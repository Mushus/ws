<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li><nuxt-link to="/tenants">入居者の管理</nuxt-link></li>
        <li>
          <nuxt-link :to="`/tenants/roomstatus/${$route.params.id}`">
            <template v-if="room">{{ room.name }}の入居状況</template>
            <template v-else>物件の入居状況</template>
          </nuxt-link>
        </li>
        <li>
          <template>入居履歴</template>
        </li>
      </Breadcrumb>
    </Content>
    <template v-if="room">
      <Content>
        <Subheader :title="`${room.name}の入居履歴`">
          <button class="btn btn__primary" @click="gotoStatus()">入居状況</button>
          <button v-if="editable" class="btn btn__primary" @click="editable = false">保存</button>
          <button v-if="!editable" class="btn btn__primary" @click="editable = true">編集</button>
        </Subheader>
      </Content>
      <Content>
        <InsertButton v-if="editable" @click="addTenant()" />
        <transition-group name="tenant" tag="div" class="tenant-group">
          <div
            v-for="(tenant, index) in formatedTenants"
            :key="tenant.raw.id"
            :class="{ tenant: true, living: tenant.isLiving }"
            >
            <div v-if="editable" class="tenant--ctrl">
              <button class="btn" @click="removeTenant(index)">
                <i class="material-icons md-18">delete</i>
              </button>
            </div>
            <TenantForm v-model="tenant.raw" :editable="editable" />
          </div>
        </transition-group>
      </Content>
    </template>
  </Container>
</template>

<script lang="ts">
import { Component, Vue } from "nuxt-property-decorator"
import { Action, Getter } from "vuex-class"
import TenantForm from '~/components/TenantForm.vue';
import { RoomShowStatus, RoomStatus } from "~/declare/article";
import * as moment from 'moment';

@Component({
  data: () => ({
    key: 0,
    now: +new Date(),
    editable: false,
    room: null,
    RoomStatus: RoomStatus,
  }),
  components: {
    TenantForm,
  }
})
export default class extends Vue {
  room: RoomShowStatus | null = null;
  now: number = 0;
  key: number = 0;
  @Action('setTitle', { namespace: 'app' }) setTitle

  created() {
    this.setTitle('入居履歴')
    // this.updateRoomStatus(+this.$route.params.id);
    this.updateRoomStatus(1);
  }

  async updateRoomStatus(id: number) {
    const data = await this.$axios.$get(`./api/rooms/status/${id}_all.json`);
    this.room = data;
  }

  gotoStatus() {
    // const id = +this.$route.params.id;
    const id = 1;
    this.$router.push(`/tenants/roomstatus/${id}`);
  }

  addTenant() {
    if (this.room == null) return;
    // マイナスの値は新規で作ったID
    this.room.tenants.unshift({
      id: --this.key,
      name: '',
      rent: this.room.rent,
      moveInAt: "",
      moveOutAt: "",
    });
  }

  removeTenant(index) {
    if (this.room == null) return;
    const showWarnning = this.room.tenants[index].id > 0;
    if (!showWarnning || confirm("この部屋を削除すると入居者の支払い履歴が一緒に削除される可能性があります。\n削除してよろしいですか？")) {
      this.room.tenants.splice(index, 1);
    }
  }

  get formatedTenants() {
    if (this.room == null) return [];
    return [...this.room.tenants].
      map(v => {
        const moveInAt = v.moveInAt == "" ? 0 : +moment(v.moveInAt);
        const moveOutAt = v.moveOutAt == "" ? 0 : +moment(v.moveOutAt);
        return {
          raw: v,
          moveInAt: moveInAt,
          moveOutAt: moveOutAt,
          isLiving: moveInAt < this.now && this.now < moveOutAt
        }
      });
  }
}
</script>

<style lang="scss" scoped>
.tenant {
  margin-bottom: 20px;
  width: 100%;
  box-sizing: border-box;
  border: 2px solid #eee;
  padding: 25px;
  position: relative;
  transition: transform .1s, opacity .1s;

  &:last-child {
    margin-bottom: 0;
  }
  &.living {
    border-color: #009688;
  }

  .tenant--ctrl {
    position: absolute;
    top: 0;
    right: 0;
  }
}
.tenant-group {
  position: relative;
}
.tenant-enter-active, .tenant-leave-active {
  position: absolute;
  transition: transform .1s, opacity .1s;
}
.tenant-move {
  transition: transform .1s;
}
.tenant-enter {
  opacity: 0;
  transform: translateY(-50px);
}
.tenant-leave-to {
  opacity: 0;
  transform: scale(0.8);
}
.tenant-leave-active {
  position: absolute;
}
</style>
