// k6 run test/loadtest.js --out=web-dashboard
// or
// curl -X POST -F image=@./test/testdata/1.4MB.jpg -o test/testdata/out.webp http://127.0.0.1:1323/v1/compress

import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  stages: [
    { duration: '15m', target: 15 },
    { duration: '10m', target: 15 },
    { duration: '15m', target: 0 },
  ],
};

const fileData = open('testdata/1.4MB.jpg', 'b'); // 実際のファイルパスに変更
const form = {
  image: http.file(fileData, 'test-image.jpg', 'image/jpeg'),
};

export default function() {
  let res = http.post('http://127.0.0.1.nip.io/v1/compress', form);
  check(res, { "status is 200": (res) => res.status === 200 });
}
