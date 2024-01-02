var x = 1

function *f1() {
    x++
    yield
    console.log('f1 x:', x)
}

function g1() {
    x++
    console.log('g1 x:', x)
}

var it = f1()
it.next()
console.log("before yield", x)
g1()
it.next()
