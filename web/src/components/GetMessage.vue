<template>
    <div class="form">
        <q-input standout v-model="getForm.msgID" :label="t('get.msgID')" label-color="yellow-3"
            :placeholder="t('get.msgIDPlaceholder')" bg-color="deep-orange-6" class="form-input" maxlength="12" />

        <q-checkbox v-model="getForm.hasMsgKey" :label="t('get.hasMsgKey')" color="orange"
            @update:model-value="getForm.msgKey = ''" />
        <q-checkbox v-model="getForm.hasSystemKey" :label="t('get.hasSystemKey')" color="red"
            @update:model-value="getForm.systemKey = ''" />

        <q-input v-show="getForm.hasMsgKey" standout v-model="getForm.msgKey" :label="t('get.msgKey')"
            label-color="cyan-12" :placeholder="t('get.msgKeyPlaceholder')" bg-color="indigo-11"
            class="form-input animate__animated animate__fadeInUp" maxlength="8" />

        <q-input v-show="getForm.hasSystemKey" standout v-model="getForm.systemKey" :label="t('get.systemKey')"
            label-color="cyan-12" :placeholder="t('get.systemKeyPlaceholder')" bg-color="orange-4"
            class="form-input animate__animated animate__fadeInUp" maxlength="20" />

        <q-btn color="green" icon="mail" :label="t('common.getMsg')" class="form-buton" @click="getMsg()" />

        <q-dialog v-model="showDialog" persistent transition-show="scale" transition-hide="scale">
            <q-card class="bg-teal text-white" style="width: 600px;max-width: 95%;">
                <q-card-section>
                    <div class="text-h6">{{ t('get.message') }}</div>
                </q-card-section>
                <q-separator />
                <q-card-section style="max-height: 50vh" class="scroll">
                    <p v-text="msg"></p>
                </q-card-section>
                <q-separator />
                <q-card-actions align="right" class="bg-white text-teal">
                    <q-btn flat :label="t('common.copy')" @click="copyMessage()" />
                    <q-btn flat :label="t('common.close')" v-close-popup />
                </q-card-actions>
            </q-card>
        </q-dialog>
    </div>
</template>
<script setup lang="ts">
import { QInput, QCheckbox, QBtn, copyToClipboard, Notify } from 'quasar';
import { useI18n } from 'vue-i18n'
import { reactive, ref } from 'vue';
import { getUserMessage } from '@/api/api';
const { t } = useI18n() // use as global scope
const showDialog = ref(false)
const getForm = reactive({
    'hasMsgKey': false,
    'hasSystemKey': false,
    'msgID': '',
    'msgKey': '',
    'systemKey': '',
})
const msg = ref('')
const getMsg = async () => {
    try {
        const response = await getUserMessage(getForm.msgID, getForm.systemKey, getForm.msgKey, false);
        msg.value = response.data.data
        showDialog.value = true
    } catch (e) { }
}
const copyMessage = () => {
    copyToClipboard(msg.value)
        .then(() => {
            Notify.create({
                color: 'primary',
                message: t('common.copySuccess'),
                icon: 'tag_faces',
                position: 'center',
                timeout: 1000
            },)
        })
        .catch(() => {
            Notify.create({
                type: 'negative',
                message: t('common.copyFail'),
                position: 'center',
                timeout: 1000
            },)
        })
}
</script>

<style scoped>
.form {
    margin: 10px auto;
    min-width: 400px;
    max-width: 60%;
}

.form-input {
    margin-top: 10px;
    margin-bottom: 14px;
}

.form-buton {
    padding: 0 10px;
    margin-left: 38%;
    margin-top: 10px;
}
</style>