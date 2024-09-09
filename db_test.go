package db_test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	tc "github.com/testcontainers/testcontainers-go/modules/compose"
)

func TestMain(m *testing.M) {
	compose, err := tc.NewDockerCompose("docker-compose.test.yaml")
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	downFunc := func() error {
		return compose.Down(context.Background(), tc.RemoveOrphans(true), tc.RemoveImagesLocal)
	}

	ctx, cancel := context.WithCancel(context.Background())

	// コンテナの起動
	if err = compose.Up(ctx, tc.Wait(true)); err != nil {
		fmt.Println(err.Error())
		cancel()
		_ = downFunc() // 手抜き
		log.Fatal(err)
	}

	code := m.Run()

	// 後処理
	cancel()
	if err = downFunc(); err != nil {
		log.Fatal(err)
	}

	os.Exit(code)
}

func TestDB(t *testing.T) {
	t.Log("Start test!!")
}
