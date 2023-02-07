<template>
  <div class="main">

    <q-card flat bordered>
      <q-card-section>
        <q-tabs inline-label class="bg-primary text-white shadow-2" v-model="tab">
          <q-tab name="get" icon="markunread_mailbox" :label="t('common.get')" class="tabbar" />
          <q-tab name="save" icon="open_in_browser" :label="t('common.save')" class="tabbar" />
          <q-tab name="edit" icon="create" :label="t('common.edit')" class="tabbar" />
        </q-tabs>
        <q-tab-panels v-model="tab" animated class="bg-purple-1 text-black">

          <q-tab-panel name="get">
            <GetMessage />
          </q-tab-panel>

          <q-tab-panel name="save">
            <SaveMessage />
          </q-tab-panel>

          <q-tab-panel name="edit">
            <EditMessage />
          </q-tab-panel>

        </q-tab-panels>

        <q-card-actions class="justify-around q-px-md">
          <q-btn flat round color="red" icon="favorite" />
          <q-btn flat round color="accent" icon="bookmark" />
          <q-btn flat round color="primary" icon="share" />
        </q-card-actions>
      </q-card-section>
    </q-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n'
import GetMessage from '@/components/GetMessage.vue';
import SaveMessage from '@/components/SaveMessage.vue';
import EditMessage from '@/components/EditMessage.vue';

const { t } = useI18n() // use as global scope
const tab = ref('get')

</script>
<style scoped>
* {
  margin: 0;
  padding: 0;
}

.main {
  margin: 3% auto;
  min-width: 500px;
  max-width: 60%;
}

.tabbar {
  padding: 10px 20px;
}
</style>
