import http from 'k6/http'
import {check , sleep} from 'k6'


export const options = {
    vus : 10,
    duration: '30s',
    thresholds: {
        http_req_duration: ['p(95) < 200'],
        http_req_failed: ['rate < 0.01'],
    },
};


export default function(){
    const res = http.get("http://localhost:8081/")
    check(
        res,
        {
            'status is 200': (r) => r.status === 200,
            'body contains Hello' : (r) => r.body.includes("Helo"),
        }
    );
    sleep(1)
} 