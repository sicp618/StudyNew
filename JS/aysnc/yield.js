const { request } = require("http");

var x = 1;

function* f1() {
  x++;
  yield;
  console.log("f1 x:", x);
}

function g1() {
  x++;
  console.log("g1 x:", x);
}

var it = f1();
it.next();
console.log("before yield", x);
g1();
it.next();

// yield 传值
function* f2(x) {
  var y = x * (yield "hello");
  return y;
}
var it2 = f2(6);
console.log(it2.next());
console.log(it2.next(7));

var a = [1, 2, 3];
var ita = a[Symbol.iterator]();
console.log(ita.next());
console.log(ita.next());
console.log(ita.next());
console.log(ita.next());

// 使用 yield* 代理迭代器
console.log("\n // 使用 yield* 代理迭代器");
function* doing() {
  try {
    var nextVal1 = 0;
    var nextVal2;
    while (true) {
      if (nextVal2 === undefined) {
        nextVal2 = 1;
      } else {
        [nextVal1, nextVal2] = [nextVal2, nextVal1 + nextVal2];
      }
      yield nextVal2;
    }
  } finally {
    console.log("clean up");
  }
}
for (var v of doing()) {
  if (v > 100) {
    break;
  }
  console.log(v);
}

// 使用 yield* 解决异步问题
const url = "http://jsonplaceholder.typicode.com"
function foo(x) {
    return fetch(url + x)
}
function *main() {
    try {
        var text = yield foo("/posts/1")
        console.log(text)
    } catch (err) {
        console.error(err)
    }
}
var it = main()
var p = it.next().value
p.then(text => {
    it.next(text)
}, err => {
    it.throw(err)
})