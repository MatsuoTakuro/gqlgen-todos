# gqlgen-todos

- [gqlgenのGetting Started](https://gqlgen.com/getting-started/)
  - [カスタムモデルを用いた暗黙的な生成](https://gqlgen.com/getting-started/#dont-eagerly-fetch-the-user)における注意点
    - 指定パッケージとの`autobind`を有効化すると、指定パッケージ配下のカスタムモデルを反映できる（上書き）
    - https://github.com/MatsuoTakuro/gqlgen-todos/blob/main/gqlgen.yml#L38

- 参考記事
  - [【GraphQL × Go】gqlgenの基本構成とオーバーフェッチを防ぐmodel resolverの実装](https://tech.layerx.co.jp/entry/2021/10/22/171242)
    - (誤表記も多少あり？)
