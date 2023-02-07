import enUS from "./en-US";
import zhCN from "./zh-CN";

import { createI18n } from 'vue-i18n';

let locale = "en-US";
if (navigator.languages != undefined)
    locale = navigator.languages[0];

const i18n = createI18n({
    locale,
    messages: {
        'en-US': enUS,
        'zh-CN': zhCN,
    },
    fallbackLocale: "en-US",
    legacy: false,
});

export default i18n;