<template>
  <Container>
    <Content>
      <Breadcrumb>
        <li><nuxt-link to="/tenants">入居者の管理</nuxt-link></li>
        <li>
          <template v-if="room">{{ room.name }}の入居状況</template>
          <template v-else>物件の入居状況</template>
        </li>
      </Breadcrumb>
    </Content>
    <template v-if="room">
      <Content>
        <Subheader :title="`${room.name}の入居状況`">
          <button class="btn btn__primary" @click="gotoHistory()">入居履歴</button>
          <button v-if="formatedTenants.length > 0" class="btn btn__primary" @click="moveOut()">退去</button>
          <template v-else>
            <button v-if="!editable" class="btn btn__primary" @click="createNew()">入居</button>
            <button v-else class="btn btn__primary" @click="saveTenant()">保存</button>
          </template>
        </Subheader>
      </Content>
      <Content>
        部屋の状況:
        <span class="badge-tag badge-tag__green badge-tag__large" v-if="room.status == RoomStatus.Living">入居中</span>
        <span class="badge-tag badge-tag__yellow badge-tag__large" v-else-if="room.status == RoomStatus.Booking">予約</span>
        <span class="badge-tag badge-tag__red badge-tag__large" v-else-if="room.status == RoomStatus.Empty">空き</span>
      </Content>
      <Content>
        <div class="tenant-group">
          <div
            v-for="tenant in formatedTenants"
            :key="tenant.raw.id"
            :class="{ tenant: true, living: tenant.isLiving }"
            >
            <TenantForm v-model="tenant.raw" :editable="editable"/>
          </div>
        </div>
        <div v-if="newTenant" :class="{ tenant: true, living: formatedNewTenant.isLiving }">
          <TenantForm v-model="formatedNewTenant.raw" :editable="editable"/>
        </div>
      </Content>
    </template>
  </Container>
</template>

<script lang="ts">
import * as moment from 'moment';
import { Component, Vue } from "nuxt-property-decorator"
import { Action, Getter } from "vuex-class"
import TenantForm from '~/components/TenantForm.vue';
import { RoomShowStatus, RoomStatus, Tenant } from "~/declare/article";

@Component({
  data: () => ({
    key: 0,
    now: +new Date(),
    editable: false,
    room: null,
    newTenant: null,
    RoomStatus: RoomStatus,
  }),
  components: {
    TenantForm,
  }
})
export default class extends Vue {
  room: RoomShowStatus | null = null;
  newTenant: Tenant | null = null;
  now: number = 0;
  key: number = 0;
  editable: boolean = false;
  @Action('setTitle', { namespace: 'app' }) setTitle

  created() {
    this.setTitle('入居状況')
    // this.updateRoomStatus(+this.$route.params.id);
    this.updateRoomStatus(1);
  }

  async updateRoomStatus(id: number) {
    const data = await this.$axios.$get(`./api/rooms/status/${id}.json`);
    this.room = data;
  }

  gotoHistory() {
    // const id = +this.$route.params.id;
    const id = 1;
    this.$router.push(`/tenants/roomstatus/history/${id}`);
  }

  get formatedTenants() {
    if (this.room == null) return [];
    return [...this.room.tenants].map(this.formatTenant);
  }

  get formatedNewTenant() {
    if (this.newTenant == null) return [];
    return this.formatTenant(this.newTenant);
  }

  formatTenant(tenant: Tenant) {
    const moveInAt = tenant.moveInAt == "" ? 0 : +moment(tenant.moveInAt);
    const moveOutAt = tenant.moveOutAt == "" ? 0 : +moment(tenant.moveOutAt);
    return {
      raw: tenant,
      moveInAt: moveInAt,
      moveOutAt: moveOutAt,
      isLiving: moveInAt < this.now && this.now < moveOutAt
    }
  }

  moveOut() {
    if (this.room == null) return;
    if (confirm('この入居者を退去させますか？')) {
      this.room.tenants = [];
    }
  }

  createNew() {
    if (this.room == null) return;
    this.newTenant = {
      id: 0,
      name: "",
      rent: this.room.rent,
      moveInAt: this.today,
      moveOutAt: this.forever,
    };
    this.editable = true;
  }

  saveTenant() {
    if (this.room == null || this.newTenant == null) return;
    this.room.tenants = [this.newTenant];
    this.newTenant = null;
    this.editable = false;
  }

  get today() {
    const now = new Date();
    return moment({
      y: now.getFullYear(),
      M: now.getMonth(),
      d: now.getDate(),
      h: 0,
      m: 0,
      s: 0,
      ms: 0,
    }).utcOffset(now.getTimezoneOffset()).toISOString();
  }

  get forever() {
    const now = new Date();
    return moment({
      y: 2999,
      M: 1,
      d: 1,
      h: 0,
      m: 0,
      s: 0,
      ms: 0,
    }).utcOffset(now.getTimezoneOffset()).toISOString();
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
</style>
