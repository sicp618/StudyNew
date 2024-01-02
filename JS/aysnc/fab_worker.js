const { parentPort } = require("worker_threads");

parentPort.on("message", function (n) {
  var result = fab(n);
  parentPort.postMessage(result);
});

function fab(n) {
  if (n < 2) {
    return 1;
  }
  return fab(n - 1) + fab(n - 2);
}
