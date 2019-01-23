# Overview

指定されたCSVの1列目をハッシュ化します。

# Installation

go version 1.2 以上 をサポートしています。

インストールしていない場合 [go install](https://golang.org/doc/install)からインストールしてください

インストールした後は、以下のコマンドから、インストールできます。
   
```bash
$ go get github.com/Suwato/dohash
```

##  CLIとしてのInstallation
pathを通して installコマンドを利用すると、cliとして利用できます。
```bash
export PATH=$PATH:$GOPATH/bin
```

* install
```bash
$ go install
```

* 実行例
```bash
$ dohash -f test.csv -a sha512 --salt dasldk
```
# Examples

```bash
dohash -f test.csv -a sha512 --salt dasldk
```

## Global Options
```
--file value, -f value
       ハッシュ化したいcsvのpathを指定してください。
```
```
--algorithm value, -a value  
        ハッシュ化のアルゴリズムを指定してください。sha256 と sha512 に対応しています。 (default: "sha256")
```
```
--stretching value           
        ストレッチングの回数を指定してください。 (default: "10")
```
```
--salt value                 
        saltを指定してください。
```
```
--help, -h                   
        show help
```
```
--version, -v                
        print the version
```

# TODO
* 進捗表示
* 並列処理
* マッピング用のCSV出力
