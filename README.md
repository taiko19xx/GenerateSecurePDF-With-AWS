# GenerateSecurePDF-With-AWS
「AWS/サーバレスでつくる！安心安全PDF」で利用したスクリプトなどをを格納しているリポジトリです。
設定や利用の方法については書籍を参照してください。

書籍は次の場所で販売中（販売していました）です。

* 商業版( https://nextpublishing.jp/book/12164.html )
* 技術書典8 / 応援祭
* BOOTH（ https://morinomiyakono.booth.pm/items/1808193 ）

## 構造
各フォルダと章の対応表は次のとおりです。

### 商業版

| フォルダ                    | 対応章                               | 中身                                            |
| --------------------------- | ------------------------------------ | ----------------------------------------------- |
| part1/local                 | 第1章 基本的な準備をする             | ローカル動作のmain.go                           |
| part1/lambda                | 第1章 基本的な準備をする             | Lambdaで動作させるmain.go                       |
| part8                       | 第8章 Webサイトに組み込む            | HTMLファイル                                    |
| example.com(直下のファイル) | 付録A/B                              | 管理用ファイルやMakefileなど                    |
| example.com/copyobject      | 第7章 API化する                      | S3のオブジェクトコピー処理のmain.goとgo.mod     |
| example.com/encrypt         | 第6章 暗号化する                     | 暗号化処理のmain.goとgo.mod                     |
| example.com/modules         | 第2章 開発周りの準備をする           | example.com以下で利用する共通モジュール         |
| example.com/property        | 第5章 カスタムプロパティーを追加する | カスタムプロパティー追加処理のmain.goとgo.mod   |
| example.com/sign            | 第4章 文字列入りページを追加する     | 文字列入りページ追加処理のmain.goとgo.mod       |
| example.com/watermark       | 第3章 ウォーターマークを追加する     | ウォーターマーク追加処理のmain.goとgo.mod       |
| pdf                         |                                      | サンプル用PDFファイル                           |

### 同人版

| フォルダ               | 対応章                           | 中身                                        |
| ---------------------- | -------------------------------- | ------------------------------------------- |
| part1/local            | 第1章 基本的な準備をしよう       | ローカル動作のmain.go                       |
| part1/lambda           | 第1章 基本的な準備をしよう       | Lambdaで動作させるmain.go                   |
| example.com/copyobject | 第6章 API化する                  | S3のオブジェクトコピー処理のmain.goとgo.mod |
| example.com/encrypt    | 第5章 暗号化する                 | 暗号化処理のmain.goとgo.mod                 |
| example.com/modules    | 第2章 開発周りの準備をしよう     | example.com以下で利用する共通モジュール     |
| example.com/sign       | 第4章 文字列入りページを追加する | 文字列入りページ追加処理のmain.goとgo.mod   |
| example.com/watermark  | 第3章 ウォーターマークを追加する | ウォーターマーク追加処理のmain.goとgo.mod   |
| pdf                    |                                  | サンプル用PDFファイル                       |

ファイル名は、本文内「リストx.x （ファイル名）」のファイル名部分に対応しています。

## ライセンス
MIT License