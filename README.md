# pubsub-emulator-example
[Zenn記事]()の参考として書いたコードです

## 使い方

### 前提条件

ローカルで Pub/Sub エミュレータが起動しており、環境変数 `PUBSUB_EMULATOR_HOST` にエミュレータのアドレス（例: `localhost:8085`）が設定されている必要があります。

```bash
$(gcloud beta emulators pubsub env-init)
```

### 実行

以下のコマンドでプログラムを実行します。

```bash
go run main.go <command>
```

`<command>` には以下のいずれかを指定します。

*   `create`: `example` というIDで新しいスキーマを作成します。
*   `update`: `example` というIDのスキーマを更新し、新しいリビジョンを作成します。

#### スキーマの作成

```bash
go run main.go create
```

成功すると、作成されたスキーマの情報が出力されます。

#### スキーマの更新

```bash
go run main.go update
```

成功すると、更新されたスキーマの情報（新しいリビジョン）が出力されます。
