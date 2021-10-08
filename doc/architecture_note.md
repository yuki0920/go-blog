
# ドメイン駆動設計について

## 参考

### DDD
- https://github.com/nrslib/itddd
- https://github.com/bxcodec/go-clean-arch

### クリーンアーキテクチャー
- https://developers.freee.co.jp/entry/clean-architecture
- https://nrslib.com/clean-architecture/

## パターン

### ドメイン知識を表現する

- 値オブジェクト: ドメインモデルを実装したドメインオブジェクト。システム固有の値を表現するために定義されたオブジェクト。同一の属性の場合は同一と見なされる。(不変)
- エンティティ: ドメインモデルを実装したドメインオブジェクト。同一の属性であっても固有のIDによって区別される。(可変)
- ドメインサービス: 値オブジェクトやエンティティなどのドメインオブジェクト不自然になる振る舞い(インスタンスメソッドとして定義するのは不自然)を定義する層。同一ドメインだが、複数のインスタンスを参照するケースなどで利用する。

### アプリケーションを実現する

- リポジトリ: ドメインオブジェクトの永続化、再構築を担う。インターフェースの定義と、インターフェースの実装(RDBMS, NoSQL等の操作)がある。
- アプリケーションサービス: ドメインオブジェクトを強調させてユースケースを実現する
- ファクトリ:

### ドメイン知識を表現する ~発展~

- 集約:
- 仕様:

## 用語

### ドメイン

- ドメインモデル: 概念の抽出
- ドメインオブジェクト: ドメインモデルをソフトウェアで動作するモジュールとして表現したもの

### ユビキタス言語

プロジェクト内の共通言語
認識の齟齬や翻訳にコストを書けないために、ドメインエキスパートとエンジニアで同一の表現をする
エンジニアが、ドメインエキスパートの表現を使うのが好ましい
例えば、`ユーザーを登録する` と `ユーザーを新規保存する` を統一する

### 境界づけられたコンテキスト

コンテキストごとにモデリングする考え方
ドメインは同一だが、コンテキストが異なる場合には、異なるモデルとして表現する
例えば、システム上、ログインするユーザーと、記事を購読するユーザーを同一ドメインの別パッケージ、別モデルで表現する

### 依存

- オブジェクトがオブジェクトを参照すること。参照(依存)する側のモジュールから参照(依存)される側のモジュールを矢印を伸ばして表現する
- 実装クラスから抽象クラスへの依存を汎化という
- インターフェース(抽象型)へ依存すると、インターフェースを実装した具象クラスであれば実態が何であっても引き渡すことができる。つまり、付け替しやすくなる。
- 抽象型を利用して具象型へ向いていた依存の方向を抽象型に変え、依存方向を成否ょすることを依存関係の逆転の法則という
- DI(Dependency Injection)とは、上位レイヤーに中小型を満たす下位レイヤーのインスタンスを引き渡すパターン

### クリーンアーキテクチャー

ビジネスルールをカプセル化したモジュール(パッケージ)を中心に添え、依存の方向を絶対的に制御する

### アーキテクチャ

何がどこに記述されるべきかといった疑問に対する回答を明確にし、ロジックが無秩序に点在することを防ぎます。