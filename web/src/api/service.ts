import axios from 'axios'
import { Notify } from 'quasar'
import i18n from '@/locales/index'
import { STATUS_CODE } from './code';

axios.defaults.baseURL = '/m'

const service = axios.create({
    timeout: 10000,
    headers: {
        'X-Requested-With': 'XMLHttpRequest',
        'Content-Type': 'application/json; charset=UTF-8'
    },
})
service.interceptors.request.use(
    config => {
        return config;
    },
    error => {
        Promise.reject(error)
    }
)

// response interceptor
service.interceptors.response.use(
    response => {
        const res = response.data;
        if (res != null && res.code != null) {
            if (res.code == STATUS_CODE.SUCCESS) {
                return Promise.resolve(response);
            } else {
                Notify.create({
                    type: 'negative',
                    message: i18n.global.t('error.' + res.code),
                    position: 'top',
                })
                return Promise.reject(response);
            }
        }
        if (res != null) {
            return Promise.resolve(response);
        }
        return Promise.reject(response);
    },
    error => {
        Notify.create({
            type: 'negative',
            message: i18n.global.t('error.' + STATUS_CODE.NETWORK_ERROR),
            position: 'top',
        })
        return Promise.reject(error)
    }
)
export default service