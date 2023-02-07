<template>
    <div class="form">

        <q-input v-model="saveForm.msg" filled clearable outlined type="textarea" color="red-12" bg-color="info"
            :label="t('save.msg')" :placeholder="t('save.msgPlaceholder')" counter maxlength="10240" />

        <q-checkbox v-model="saveForm.hasMsgKey" :label="t('save.hasMsgKey')" color="orange"
            @update:model-value="saveForm.msgKey = ''" />
        <q-checkbox v-model="saveForm.hasSystemKey" :label="t('save.hasSystemKey')" color="red"
            @update:model-value="saveForm.systemKey = ''" />

        <q-input v-show="saveForm.hasMsgKey" standout v-model="saveForm.msgKey" :label="t('save.msgKey')"
            label-color="cyan-12" :placeholder="t('save.msgKeyPlaceholder')" bg-color="indigo-11"
            class="form-input animate__animated animate__fadeInUp" maxlength="8" />

        <q-input v-show="saveForm.hasSystemKey" standout v-model="saveForm.systemKey" :label="t('save.systemKey')"
            label-color="cyan-12" :placeholder="t('save.systemKeyPlaceholder')" bg-color="orange-4"
            class="form-input animate__animated animate__fadeInUp" maxlength="20" />

        <q-input standout v-model="saveForm.editKey" :label="t('save.editKey')" label-color="yellow-3"
            :placeholder="t('save.editKeyPlaceholder')" bg-color="deep-orange-6" class="form-input" maxlength="8" />

        <q-btn color="red" icon="send" :label="t('common.saveMsg')" class="form-buton" @click="saveMsg()" />

        <q-dialog v-model="showDialog" persistent transition-show="scale" transition-hide="scale">
            <q-card class="bg-teal text-white" style="width: 600px;max-width: 95%;">
                <q-card-section>
                    <div class="text-h6">{{ t('common.alert') }}</div>
                </q-card-section>
                <q-separator />
                <q-card-section style="max-height: 50vh" class="scroll">
                    <p v-html="t('save.saveSuccess', { msgID })"></p>
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
import { saveUserMessage } from '@/api/api';
import { QInput, QCheckbox, QBtn, copyToClipboard, Notify, QCard, QCardActions, QCardSection, QDialog, QSeparator } from 'quasar';
import { useI18n } from 'vue-i18n'
import { reactive, ref } from 'vue';
const { t } = useI18n() // use as global scope
const showDialog = ref(false)
const saveForm = reactive({
    'hasMsgKey': true,
    'hasSystemKey': false,
    'editKey': '',
    'msgKey': '',
    'systemKey': '',
    'msg': '',
})

const msgID = ref('')

const saveMsg = async () => {
    try {
        const response = await saveUserMessage(saveForm.msg, saveForm.msgKey, saveForm.systemKey, saveForm.editKey)
        msgID.value = response.data.data
        showDialog.value = true;
    } catch (e) { }
}

const copyMessage = () => {
    copyToClipboard(msgID.value)
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