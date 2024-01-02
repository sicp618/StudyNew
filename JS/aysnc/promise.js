//
var p = Promise.resolve(21)
p.then(function(v) {
    console.log(v)
    return v * 2
}).then(function(v) {
    console.log(v)
})

var p1 = Promise.resolve(32)
p1.then(function(v) {
    console.log(v)
    return new Promise(function(resolve, reject) {
        setTimeout(function() {
            resolve(v * 2)
        }, 1000)
    })
})
.then(function(v) {
    console.log(v)
})


function delay(time) {
    return new Promise(function(resolve, reject) {
        setTimeout(resolve, time)
    })
}

delay(1000).then(function STEP2() {
    console.log("step 2 (after 100ms)")
    return delay(1000)
}).then(function STEP3() {
    foo.bar()
    console.log("step 3 (after another 100ms)")
})
.then(function STEP4() {
    console.log("step 4 (next Job)")
    return delay(1000)
}, function reject(err) {
    console.log("reject", err)
}).then(function STEP5() {
    console.log("step 5 (after another 100ms)")
})