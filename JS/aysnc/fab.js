const {
  Worker,
  isMainThread,
  parentPort,
  workerData,
} = require("worker_threads");

var worker = new Worker("./fab_worker.js");
worker.postMessage(10);
worker.on("message", (msg) => {
  console.log("fab 10:", msg);
});

setTimeout(() => {
  worker.terminate();
}, 2000);
