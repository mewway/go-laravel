package filesystem

import (
	"context"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/gookit/color"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/ory/dockertest/v3"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	configmocks "github.com/goravel/framework/contracts/config/mocks"
	"github.com/goravel/framework/contracts/filesystem"
	"github.com/goravel/framework/support/file"
	supporttime "github.com/goravel/framework/support/time"
	testingdocker "github.com/goravel/framework/testing/docker"
)

type TestDisk struct {
	disk string
	url  string
}

func TestStorage(t *testing.T) {
	if !file.Exists("../.env") && os.Getenv("AWS_ACCESS_KEY_ID") == "" {
		color.Redln("No filesystem tests run, need create .env based on .env.example, then initialize it")
		return
	}

	assert.Nil(t, file.Create("test.txt", "Goravel"))
	mockConfig := initConfig()
	minioPool, minioResource, err := initMinioDocker(mockConfig)
	assert.Nil(t, err)

	var driver filesystem.Driver

	disks := []TestDisk{
		{
			disk: "local",
			url:  "http://localhost/storage",
		},
		{
			disk: "oss",
			url:  mockConfig.GetString("filesystems.disks.oss.url"),
		},
		{
			disk: "cos",
			url:  mockConfig.GetString("filesystems.disks.cos.url"),
		},
		{
			disk: "s3",
			url:  mockConfig.GetString("filesystems.disks.s3.url"),
		},
		{
			disk: "minio",
			url:  mockConfig.GetString("filesystems.disks.minio.url"),
		},
		{
			disk: "custom",
			url:  "http://localhost/storage",
		},
	}

	tests := []struct {
		name  string
		setup func(disk TestDisk)
	}{
		{
			name: "AllDirectories",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("AllDirectories/1.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("AllDirectories/2.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("AllDirectories/3/3.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("AllDirectories/3/5/6/6.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.MakeDirectory("AllDirectories/3/4"), disk.disk)
				assert.True(t, driver.Exists("AllDirectories/1.txt"), disk.disk)
				assert.True(t, driver.Exists("AllDirectories/2.txt"), disk.disk)
				assert.True(t, driver.Exists("AllDirectories/3/3.txt"), disk.disk)
				assert.True(t, driver.Exists("AllDirectories/3/4/"), disk.disk)
				assert.True(t, driver.Exists("AllDirectories/3/5/6/6.txt"), disk.disk)
				files, err := driver.AllDirectories("AllDirectories")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/", "3/4/", "3/5/", "3/5/6/"}, files)
				files, err = driver.AllDirectories("./AllDirectories")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/", "3/4/", "3/5/", "3/5/6/"}, files)
				files, err = driver.AllDirectories("/AllDirectories")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/", "3/4/", "3/5/", "3/5/6/"}, files)
				files, err = driver.AllDirectories("./AllDirectories/")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/", "3/4/", "3/5/", "3/5/6/"}, files)
				assert.Nil(t, driver.DeleteDirectory("AllDirectories"), disk.disk)
			},
		},
		{
			name: "AllFiles",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("AllFiles/1.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("AllFiles/2.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("AllFiles/3/3.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("AllFiles/3/4/4.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("AllFiles/1.txt"), disk.disk)
				assert.True(t, driver.Exists("AllFiles/2.txt"), disk.disk)
				assert.True(t, driver.Exists("AllFiles/3/3.txt"), disk.disk)
				assert.True(t, driver.Exists("AllFiles/3/4/4.txt"), disk.disk)
				files, err := driver.AllFiles("AllFiles")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt", "3/3.txt", "3/4/4.txt"}, files)
				files, err = driver.AllFiles("./AllFiles")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt", "3/3.txt", "3/4/4.txt"}, files)
				files, err = driver.AllFiles("/AllFiles")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt", "3/3.txt", "3/4/4.txt"}, files)
				files, err = driver.AllFiles("./AllFiles/")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt", "3/3.txt", "3/4/4.txt"}, files)
				assert.Nil(t, driver.DeleteDirectory("AllFiles"), disk.disk)
			},
		},
		{
			name: "Copy",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Copy/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Copy/1.txt"), disk.disk)
				assert.Nil(t, driver.Copy("Copy/1.txt", "Copy1/1.txt"), disk.disk)
				assert.True(t, driver.Exists("Copy/1.txt"), disk.disk)
				assert.True(t, driver.Exists("Copy1/1.txt"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("Copy"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("Copy1"), disk.disk)
			},
		},
		{
			name: "Delete",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Delete/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Delete/1.txt"), disk.disk)
				assert.Nil(t, driver.Delete("Delete/1.txt"), disk.disk)
				assert.True(t, driver.Missing("Delete/1.txt"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("Delete"), disk.disk)
			},
		},
		{
			name: "DeleteDirectory",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("DeleteDirectory/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("DeleteDirectory/1.txt"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("DeleteDirectory"), disk.disk)
				assert.True(t, driver.Missing("DeleteDirectory/1.txt"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("DeleteDirectory"), disk.disk)
			},
		},
		{
			name: "Directories",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Directories/1.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("Directories/2.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("Directories/3/3.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("Directories/3/5/5.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.MakeDirectory("Directories/3/4"), disk.disk)
				assert.True(t, driver.Exists("Directories/1.txt"), disk.disk)
				assert.True(t, driver.Exists("Directories/2.txt"), disk.disk)
				assert.True(t, driver.Exists("Directories/3/3.txt"), disk.disk)
				assert.True(t, driver.Exists("Directories/3/4/"), disk.disk)
				assert.True(t, driver.Exists("Directories/3/5/5.txt"), disk.disk)
				files, err := driver.Directories("Directories")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/"}, files)
				files, err = driver.Directories("./Directories")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/"}, files)
				files, err = driver.Directories("/Directories")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/"}, files)
				files, err = driver.Directories("./Directories/")
				assert.Nil(t, err)
				assert.Equal(t, []string{"3/"}, files)
				assert.Nil(t, driver.DeleteDirectory("Directories"), disk.disk)
			},
		},
		{
			name: "Files",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Files/1.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("Files/2.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("Files/3/3.txt", "Goravel"), disk.disk)
				assert.Nil(t, driver.Put("Files/3/4/4.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Files/1.txt"), disk.disk)
				assert.True(t, driver.Exists("Files/2.txt"), disk.disk)
				assert.True(t, driver.Exists("Files/3/3.txt"), disk.disk)
				assert.True(t, driver.Exists("Files/3/4/4.txt"), disk.disk)
				files, err := driver.Files("Files")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt"}, files)
				files, err = driver.Files("./Files")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt"}, files)
				files, err = driver.Files("/Files")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt"}, files)
				files, err = driver.Files("./Files/")
				assert.Nil(t, err)
				assert.Equal(t, []string{"1.txt", "2.txt"}, files)
				assert.Nil(t, driver.DeleteDirectory("Files"), disk.disk)
			},
		},
		{
			name: "Get",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Get/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Get/1.txt"), disk.disk)
				data, err := driver.Get("Get/1.txt")
				assert.Nil(t, err)
				assert.Equal(t, "Goravel", data)
				length, err := driver.Size("Get/1.txt")
				assert.Nil(t, err)
				assert.Equal(t, int64(7), length)
				assert.Nil(t, driver.DeleteDirectory("Get"), disk.disk)
			},
		},
		{
			name: "LastModified",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("LastModified/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("LastModified/1.txt"), disk.disk)
				date, err := driver.LastModified("LastModified/1.txt")
				assert.Nil(t, err)

				l, err := time.LoadLocation("UTC")
				assert.Nil(t, err, disk.disk)
				assert.Equal(t, time.Now().In(l).Format("2006-01-02 15"), date.Format("2006-01-02 15"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("LastModified"), disk.disk)
			},
		},
		{
			name: "MakeDirectory",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.MakeDirectory("MakeDirectory1/"), disk.disk)
				assert.Nil(t, driver.MakeDirectory("MakeDirectory2"), disk.disk)
				assert.Nil(t, driver.MakeDirectory("MakeDirectory3/MakeDirectory4"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("MakeDirectory1"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("MakeDirectory2"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("MakeDirectory3"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("MakeDirectory4"), disk.disk)
			},
		},
		{
			name: "MimeType",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("MimeType/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("MimeType/1.txt"), disk.disk)
				mimeType, err := driver.MimeType("MimeType/1.txt")
				assert.Nil(t, err, disk.disk)
				mediaType, _, err := mime.ParseMediaType(mimeType)
				assert.Nil(t, err, disk.disk)
				assert.Equal(t, "text/plain", mediaType, disk.disk)

				fileInfo, err := NewFile("../logo.png")
				assert.Nil(t, err, disk.disk)
				path, err := driver.PutFile("MimeType", fileInfo)
				assert.Nil(t, err, disk.disk)
				assert.True(t, driver.Exists(path), disk.disk)
				mimeType, err = driver.MimeType(path)
				assert.Nil(t, err, disk.disk)
				assert.Equal(t, "image/png", mimeType, disk.disk)
			},
		},
		{
			name: "Move",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Move/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Move/1.txt"), disk.disk)
				assert.Nil(t, driver.Move("Move/1.txt", "Move1/1.txt"), disk.disk)
				assert.True(t, driver.Missing("Move/1.txt"), disk.disk)
				assert.True(t, driver.Exists("Move1/1.txt"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("Move"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("Move1"), disk.disk)
			},
		},
		{
			name: "Put",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Put/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Put/1.txt"), disk.disk)
				assert.True(t, driver.Missing("Put/2.txt"), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("Put"), disk.disk)
			},
		},
		{
			name: "PutFile_Image",
			setup: func(disk TestDisk) {
				fileInfo, err := NewFile("../logo.png")
				assert.Nil(t, err)
				path, err := driver.PutFile("PutFile1", fileInfo)
				assert.Nil(t, err)
				assert.True(t, driver.Exists(path), disk.disk)
				assert.Nil(t, driver.DeleteDirectory("PutFile1"), disk.disk)
			},
		},
		{
			name: "PutFile_Text",
			setup: func(disk TestDisk) {
				fileInfo, err := NewFile("./test.txt")
				assert.Nil(t, err)
				path, err := driver.PutFile("PutFile", fileInfo)
				assert.Nil(t, err)
				assert.True(t, driver.Exists(path), disk.disk)
				data, err := driver.Get(path)
				assert.Nil(t, err)
				assert.Equal(t, "Goravel", data)
				assert.Nil(t, driver.DeleteDirectory("PutFile"), disk.disk)
			},
		},
		{
			name: "PutFileAs_Text",
			setup: func(disk TestDisk) {
				fileInfo, err := NewFile("./test.txt")
				assert.Nil(t, err)
				path, err := driver.PutFileAs("PutFileAs", fileInfo, "text")
				assert.Nil(t, err)
				assert.Equal(t, "PutFileAs/text.txt", path)
				assert.True(t, driver.Exists(path), disk.disk)
				data, err := driver.Get(path)
				assert.Nil(t, err)
				assert.Equal(t, "Goravel", data)

				path, err = driver.PutFileAs("PutFileAs", fileInfo, "text1.txt")
				assert.Nil(t, err)
				assert.Equal(t, "PutFileAs/text1.txt", path)
				assert.True(t, driver.Exists(path), disk.disk)
				data, err = driver.Get(path)
				assert.Nil(t, err)
				assert.Equal(t, "Goravel", data)

				assert.Nil(t, driver.DeleteDirectory("PutFileAs"), disk.disk)
			},
		},
		{
			name: "PutFileAs_Image",
			setup: func(disk TestDisk) {
				fileInfo, err := NewFile("../logo.png")
				assert.Nil(t, err)
				path, err := driver.PutFileAs("PutFileAs1", fileInfo, "image")
				assert.Nil(t, err)
				assert.Equal(t, "PutFileAs1/image.png", path)
				assert.True(t, driver.Exists(path), disk.disk)

				path, err = driver.PutFileAs("PutFileAs1", fileInfo, "image1.png")
				assert.Nil(t, err)
				assert.Equal(t, "PutFileAs1/image1.png", path)
				assert.True(t, driver.Exists(path), disk.disk)

				assert.Nil(t, driver.DeleteDirectory("PutFileAs1"), disk.disk)
			},
		},
		{
			name: "Size",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Size/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Size/1.txt"), disk.disk)
				length, err := driver.Size("Size/1.txt")
				assert.Nil(t, err)
				assert.Equal(t, int64(7), length)
				assert.Nil(t, driver.DeleteDirectory("Size"), disk.disk)
			},
		},
		{
			name: "TemporaryUrl",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("TemporaryUrl/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("TemporaryUrl/1.txt"), disk.disk)
				url, err := driver.TemporaryUrl("TemporaryUrl/1.txt", supporttime.Now().Add(5*time.Second))
				assert.Nil(t, err)
				assert.NotEmpty(t, url)
				if disk.disk != "local" && disk.disk != "custom" {
					resp, err := http.Get(url)
					assert.Nil(t, err)
					content, err := ioutil.ReadAll(resp.Body)
					assert.Nil(t, resp.Body.Close())
					assert.Nil(t, err)
					assert.Equal(t, "Goravel", string(content), disk.disk)
				}
				assert.Nil(t, driver.DeleteDirectory("TemporaryUrl"), disk.disk)
			},
		},
		{
			name: "Url",
			setup: func(disk TestDisk) {
				assert.Nil(t, driver.Put("Url/1.txt", "Goravel"), disk.disk)
				assert.True(t, driver.Exists("Url/1.txt"), disk.disk)
				url := disk.url + "/Url/1.txt"
				assert.Equal(t, url, driver.Url("Url/1.txt"), disk.disk)
				if disk.disk != "local" && disk.disk != "custom" {
					resp, err := http.Get(url)
					assert.Nil(t, err)
					content, err := ioutil.ReadAll(resp.Body)
					assert.Nil(t, resp.Body.Close())
					assert.Nil(t, err)
					assert.Equal(t, "Goravel", string(content), disk.disk)
				}
				assert.Nil(t, driver.DeleteDirectory("Url"), disk.disk)
			},
		},
	}

	for _, disk := range disks {
		var err error
		driver, err = NewDriver(mockConfig, disk.disk)
		assert.NotNil(t, driver)
		assert.Nil(t, err)
		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				test.setup(disk)
			})
		}

		if disk.disk == "local" || disk.disk == "custom" {
			file.Remove("./storage")
		}
	}

	file.Remove("test.txt")
	assert.Nil(t, minioPool.Purge(minioResource))
}

func initConfig() *configmocks.Config {
	mockConfig := &configmocks.Config{}
	ConfigFacade = mockConfig
	mockConfig.On("GetString", "app.timezone").Return("UTC")
	mockConfig.On("GetString", "filesystems.default").Return("local")
	mockConfig.On("GetString", "filesystems.disks.local.driver").Return("local")
	mockConfig.On("GetString", "filesystems.disks.local.root").Return("storage/app")
	mockConfig.On("GetString", "filesystems.disks.local.url").Return("http://localhost/storage")
	mockConfig.On("GetString", "filesystems.disks.custom.driver").Return("custom")
	mockConfig.On("Get", "filesystems.disks.custom.via").Return(&Local{
		config: mockConfig,
		root:   "storage/app/public",
		url:    "http://localhost/storage",
	})
	mockConfig.On("GetString", "filesystems.disks.s3.driver").Return("s3")
	mockConfig.On("GetString", "filesystems.disks.oss.driver").Return("oss")
	mockConfig.On("GetString", "filesystems.disks.cos.driver").Return("cos")
	mockConfig.On("GetString", "filesystems.disks.minio.driver").Return("minio")
	mockConfig.On("GetString", "filesystems.disks.minio.region").Return("")
	mockConfig.On("GetBool", "filesystems.disks.minio.ssl", false).Return(false)

	if file.Exists("../.env") {
		vip := viper.New()
		vip.SetConfigName("../.env")
		vip.SetConfigType("env")
		vip.AddConfigPath(".")
		_ = vip.ReadInConfig()
		vip.SetEnvPrefix("goravel")
		vip.AutomaticEnv()

		mockConfig.On("GetString", "filesystems.disks.s3.key").Return(vip.Get("AWS_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.s3.secret").Return(vip.Get("AWS_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.s3.region").Return(vip.Get("AWS_DEFAULT_REGION"))
		mockConfig.On("GetString", "filesystems.disks.s3.bucket").Return(vip.Get("AWS_BUCKET"))
		mockConfig.On("GetString", "filesystems.disks.s3.url").Return(vip.Get("AWS_URL"))
		mockConfig.On("GetString", "filesystems.disks.oss.key").Return(vip.Get("ALIYUN_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.oss.secret").Return(vip.Get("ALIYUN_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.oss.bucket").Return(vip.Get("ALIYUN_BUCKET"))
		mockConfig.On("GetString", "filesystems.disks.oss.url").Return(vip.Get("ALIYUN_URL"))
		mockConfig.On("GetString", "filesystems.disks.oss.endpoint").Return(vip.Get("ALIYUN_ENDPOINT"))
		mockConfig.On("GetString", "filesystems.disks.cos.key").Return(vip.Get("TENCENT_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.cos.secret").Return(vip.Get("TENCENT_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.cos.bucket").Return(vip.Get("TENCENT_BUCKET"))
		mockConfig.On("GetString", "filesystems.disks.cos.url").Return(vip.Get("TENCENT_URL"))
		mockConfig.On("GetString", "filesystems.disks.minio.key").Return(vip.Get("MINIO_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.minio.secret").Return(vip.Get("MINIO_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.minio.bucket").Return(vip.Get("MINIO_BUCKET"))
	}
	if os.Getenv("AWS_ACCESS_KEY_ID") != "" {
		mockConfig.On("GetString", "filesystems.disks.s3.key").Return(os.Getenv("AWS_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.s3.secret").Return(os.Getenv("AWS_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.s3.region").Return(os.Getenv("AWS_DEFAULT_REGION"))
		mockConfig.On("GetString", "filesystems.disks.s3.bucket").Return(os.Getenv("AWS_BUCKET"))
		mockConfig.On("GetString", "filesystems.disks.s3.url").Return(os.Getenv("AWS_URL"))
		mockConfig.On("GetString", "filesystems.disks.oss.key").Return(os.Getenv("ALIYUN_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.oss.secret").Return(os.Getenv("ALIYUN_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.oss.bucket").Return(os.Getenv("ALIYUN_BUCKET"))
		mockConfig.On("GetString", "filesystems.disks.oss.url").Return(os.Getenv("ALIYUN_URL"))
		mockConfig.On("GetString", "filesystems.disks.oss.endpoint").Return(os.Getenv("ALIYUN_ENDPOINT"))
		mockConfig.On("GetString", "filesystems.disks.cos.key").Return(os.Getenv("TENCENT_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.cos.secret").Return(os.Getenv("TENCENT_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.cos.bucket").Return(os.Getenv("TENCENT_BUCKET"))
		mockConfig.On("GetString", "filesystems.disks.cos.url").Return(os.Getenv("TENCENT_URL"))
		mockConfig.On("GetString", "filesystems.disks.minio.key").Return(os.Getenv("MINIO_ACCESS_KEY_ID"))
		mockConfig.On("GetString", "filesystems.disks.minio.secret").Return(os.Getenv("MINIO_ACCESS_KEY_SECRET"))
		mockConfig.On("GetString", "filesystems.disks.minio.bucket").Return(os.Getenv("MINIO_BUCKET"))
	}

	return mockConfig
}

func initMinioDocker(mockConfig *configmocks.Config) (*dockertest.Pool, *dockertest.Resource, error) {
	pool, err := testingdocker.Pool()
	if err != nil {
		return nil, nil, err
	}

	key := mockConfig.GetString("filesystems.disks.minio.key")
	secret := mockConfig.GetString("filesystems.disks.minio.secret")
	bucket := mockConfig.GetString("filesystems.disks.minio.bucket")
	resource, err := testingdocker.Resource(pool, &dockertest.RunOptions{
		Repository: "minio/minio",
		Tag:        "latest",
		Env: []string{
			"MINIO_ACCESS_KEY=" + key,
			"MINIO_SECRET_KEY=" + secret,
		},
		Cmd: []string{
			"server",
			"/data",
		},
		ExposedPorts: []string{
			"9000/tcp",
		},
	})
	if err != nil {
		return nil, nil, err
	}

	_ = resource.Expire(600)

	endpoint := fmt.Sprintf("127.0.0.1:%s", resource.GetPort("9000/tcp"))
	mockConfig.On("GetString", "filesystems.disks.minio.url").Return(fmt.Sprintf("http://%s/%s", endpoint, bucket))
	mockConfig.On("GetString", "filesystems.disks.minio.endpoint").Return(endpoint)

	if err := pool.Retry(func() error {
		client, err := minio.New(endpoint, &minio.Options{
			Creds: credentials.NewStaticV4(key, secret, ""),
		})
		if err != nil {
			return err
		}

		if err := client.MakeBucket(context.Background(), mockConfig.GetString("filesystems.disks.minio.bucket"), minio.MakeBucketOptions{}); err != nil {
			return err
		}

		policy := `{
        "Version": "2012-10-17",
        "Statement": [
            {
                "Action": [
                    "s3:GetObject",
                    "s3:PutObject"
                ],
                "Effect": "Allow",
                "Principal": "*",
                "Resource": [
                    "arn:aws:s3:::` + bucket + `/*"
                ]
            },
            {
                "Action": [
                    "s3:ListBucket"
                ],
                "Effect": "Allow",
                "Principal": "*",
                "Resource": [
                    "arn:aws:s3:::` + bucket + `"
                ]
            }
        ]
    }`

		if err := client.SetBucketPolicy(context.Background(), bucket, policy); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, nil, err
	}

	return pool, resource, nil
}
