// [必須] 積分区間
var interval = [0, Math.PI];

// [必須] 標本数
var n = 1000;

// [必須] 積分対象の関数
function f(x) {
    return Math.sin(x);
}

// [オプション] エラー率を算出するための正しい積分結果の指定 
var answer = 2;
