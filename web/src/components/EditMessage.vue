<template>
    <div class="form">
        <q-input standout v-model="editForm.msgID" :label="t('edit.msgID')" label-color="yellow-3"
            :placeholder="t('edit.msgIDPlaceholder')" bg-color="deep-orange-6" class="form-input" maxlength="12" />

        <q-input standout v-model="editForm.editKey" :label="t('edit.editKey')" label-color="yellow-3"
            :placeholder="t('edit.editKeyPlaceholder')" bg-color="deep-orange-6" class="form-input" maxlength="8" />

        <q-checkbox v-model="editForm.hasSystemKey" :label="t('edit.hasSystemKey')" color="red"
            @update:model-value="editForm.systemKey = ''" />

        <q-input v-show="editForm.hasSystemKey" standout v-model="editForm.systemKey" :label="t('edit.systemKey')"
            label-color="cyan-12" :placeholder="t('edit.systemKeyPlaceholder')" bg-color="orange-4"
            class="form-input animate__animated animate__fadeInUp" maxlength="20" />

        <q-input v-model="editForm.msg" filled clearable outlined type="textarea" color="red-12" bg-color="info"
            :label="t('edit.msg')" :placeholder="t('edit.msgPlaceholder')" counter maxlength="10240" />

        <q-btn color="primary" icon="send" :label="t('common.editMsg')" class="form-buton" @click="editMsg()" />
    </div>
</template>
<script setup lang="ts">
import { editUserMessage } from '@/api/api';
import { QInput, QCheckbox, QBtn } from 'quasar';
import { reactive } from 'vue';
import { useI18n } from 'vue-i18n';
import { useQuasar } from 'quasar'
const $q = useQuasar()
const { t } = useI18n() // use as global scope
const editForm = reactive({
    'hasSystemKey': false,
    'msgID': '',
    'editKey': '',
    'systemKey': '',
    'msg': '',
})
const editMsg = async () => {
    try {
        await editUserMessage(editForm.msgID, editForm.msg, editForm.systemKey, editForm.editKey)
        $q.dialog({
            title: t('common.alert'),
            message: t('edit.editSuccess')
        })
    } catch (e) { }
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