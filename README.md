# gointegral --- golang implementation of Simpson method

シンプソン法の golang 実装。

十分小さな区間 ![](https://latex.codecogs.com/gif.latex?[x_i,x_{i+1}]) における関数 ![](https://latex.codecogs.com/gif.latex?f(x)) の定積分は、次の近似式で与えられる。

![](https://latex.codecogs.com/gif.latex?\int_{x_i}^{x_{i&plus;1}}f(x)dx\approx\frac{1}{6}(x_{i&plus;1}-x_i)(f(x_i)&plus;4f(\frac{x_i&plus;x_{i&plus;1}}{2})&plus;f(x_{i&plus;1})))

例えば区間 ![](https://latex.codecogs.com/gif.latex?[a,b]) の定積分であれば、区間を N 等分した十分に小さな区間 ![](https://latex.codecogs.com/gif.latex?[x_i,x_{i+1}]) の定積分の合計と考えればよい。

![](https://latex.codecogs.com/gif.latex?x_i=a&plus;\frac{b-a}{N}i) where ![](https://latex.codecogs.com/gif.latex?(0\leq{i}<N))

![](https://latex.codecogs.com/gif.latex?%5Cint_%7Ba%7D%5E%7Bb%7Df%28x%29dx%5Capprox%5Csum_%7Bi%3D0%7D%5E%7BN-1%7D%5Cleft%5C%7B%5Cfrac%7B1%7D%7B6%7D%28x_%7Bi&plus;1%7D-x_i%29%28f%28x_i%29&plus;4f%28%5Cfrac%7Bx_i&plus;x_%7Bi&plus;1%7D%7D%7B2%7D%29&plus;f%28x_%7Bi&plus;1%7D%29%29%5Cright%5C%7D)

<!--
![](https://latex.codecogs.com/gif.latex?\int_{a}^{b}f(x)dx\approx\sum_{i=0}^{N-1}\left\{\frac{1}{6}(x_{i&plus;1}-x_i)(f(x_i)&plus;4f(\frac{x_i&plus;x_{i&plus;1}}{2})&plus;f(x_{i&plus;1}))\right\})
-->

さて、gointegral では積分区間、区間分割数、積分対象とする関数を JavaScript のコードで与える。sin 関数の[0,1] 区間での定積分では次のような JavaScript コードを使う。

```javascript
// 積分区間 [a, b]
var interval = [0, Math.PI];

// 区間分割数 N
var n = 1000;

// 積分の対象とする関数 f(x)
function f(x) {
    return Math.sin(x);
}
```

gointegral の使い方は次の通りである。 

```
$ ./gointegral sin.js
interval = [0 3.141592653589793]
n = 1000
result = 1.999995065201925
$ 
```

積分結果がわかっていて、シンプソン法のエラー率[%]を求めたい場合には、次のように answer 変数を指定する。

```javascript
// 積分区間 [a, b]
var interval = [0, Math.PI];

// 区間分割数 N
var n = 1000;

// 積分の対象とする関数 f(x)
function f(x) {
    return Math.sin(x);
}

// エラー率[%]を算出するための正しい積分結果の指定 
var answer = 2;
```

実行結果は以下のとおりである。

```
$ ./gointegral sin.js
interval = [0 3.141592653589793]
n = 1000
result = 1.999995065201925
answer = 2
error  = 0.0002467399037531237 [%]
```