// k6 run test/loadtest.js
// or
// curl -X POST -F image=@./test/testdata/1.4MB.jpg -o test/testdata/out.webp http://127.0.0.1:1323/v1/compress

import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
  stages: [
    { duration: '30s', target: 20 }, // 最初の30秒でVUsを0から20まで増加
    //{ duration: '1m', target: 10 },  // 次の1分間でVUsを20から10まで減少
    //{ duration: '30s', target: 0 },  // 最後の30秒でVUsを10から0まで減少
  ],
};

const fileData = open('test/testdata/1.4MB.jpg', 'b'); // 実際のファイルパスに変更
const form = {
  image: http.file(fileData, 'test-image.jpg', 'image/jpeg'),
};

export default function() {
  let res = http.post('http://127.0.0.1:1323/v1/compress', form);
  check(res, { "status is 200": (res) => res.status === 200 });
}
