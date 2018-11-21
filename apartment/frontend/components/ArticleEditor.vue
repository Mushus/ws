<template>
  <div>
    <template v-if="model">
      <InputItem>
        <label for="articleName">建物名</label>
        <input class="input" id="articleName" type="text" v-model="model.name" :readonly="!editable">
      </InputItem>
      <InputItem>
        <label>部屋</label>
        <transition-group name="room" tag="div" class="rooms">
          <div
            v-for="(room, index) in sortedRoom"
            :key="index"
            :class="{ room: true, room__dragging: room == draggingRoom, room__draggable: editable }"
            @dragstart="dragStartRoom(room)"
            @dragend="dragEndRoom()"
            @dragenter="dragEnterRoom(room)"
            :draggable="editable"
            >
            <div v-if="editable" class="room--ctrl">
              <button class="btn" @click="removeRoom(index)">
                <i class="material-icons md-18">delete</i>
              </button>
            </div>
            <InputItem>
              <label :for="`room${index}-roomName`">部屋名</label>
              <input class="input" :id="`room${index}-roomName`" type="text" v-model="room.name" :readonly="!editable">
            </InputItem>
            <InputItem>
              <label :for="`room${index}-rent`">家賃</label>
              <input class="input" :id="`room${index}-rent`" type="number" min="0" step="1000" v-model="room.rent" :readonly="!editable">
            </InputItem>
          </div>
        </transition-group>
      </InputItem>
      <InputItem v-if="editable">
        <InsertButton v-if="editable" @click="addRoom()" />
      </InputItem>
    </template>
  </div>
</template>

<script lang="ts">
import * as numeral from 'numeral';
import Vue from "vue";
import { Component, Prop, Watch, Emit } from "nuxt-property-decorator";
import { ArticleDetail, Room } from '~/declare/article';

@Component({
  data: () => ({
    model: null,
    draggingRoom: null,
  }),
  props: {
    value: Object,
    editable: {
      default: false,
      type: Boolean
    }
  }
})
export default class ArticleEditor extends Vue {
  model: ArticleDetail | null = null;
  draggingRoom: Room | null = null;
  value: any;

  created () {
    this.model = this.value;
  }

  @Watch("value")
  onValueUpdate(val: any, oldVal: any) {
    this.model = val;
  }

  @Watch("model", { deep: true })
  onModelUpdate(val: any, oldVal: any) {
    this.$emit('input', val)
  }

  addRoom() {
    if (this.model == null) return;
    this.model.rooms.push({
      id: 0,
      name: "",
      rent: 0,
      index: this.model.rooms.length
    });
  }

  commaNum(num: number) {
    return numeral(num).format('0,0')
  }

  removeRoom(index) {
    if (this.model == null) return;
    const showWarnning = this.model.rooms[index].id != 0;
    if (!showWarnning || confirm("この部屋を削除すると入居者情報が一緒に削除される可能性があります。\n削除してよろしいですか？")) {
      this.model.rooms.splice(index, 1);
      // 削除した分indexがずれるので振り直す
      [...this.model.rooms].sort((a, b) => a.index - b.index).forEach((room, idx) => {
        room.index = idx;
      });
    }
  }

  get sortedRoom() {
    if (this.model == null) return [];
    return [...this.model.rooms].sort((a, b) => a.index - b.index);
  }

  dragStartRoom(room: Room) {
    this.draggingRoom = room;
  }

  dragEndRoom() {
    this.draggingRoom = null;
  }

  dragEnterRoom(room: Room) {
    if (this.draggingRoom == null) return;
    const tmpIndex = room.index;
    room.index = this.draggingRoom.index;
    this.draggingRoom.index = tmpIndex;
  }
}
</script>

<style lang="scss" scoped>

label {
  font-weight: bold;
}
.rooms {
  position: relative;

  .room {
    position: relative;
    width: 100%;
    box-sizing: border-box;
    padding: 25px;
    border: 2px solid #ddd;
    margin-bottom: -2px;
    transition: transform .1s, opacity .1s, height .1s;

    &.room__dragging {
      background-color: #E3F2FD;
    }
    &.room__draggable {
      cursor: move;
      &:hover {
        z-index: 10;
        box-shadow: 0 0 3px 0 #2196F3;
      }
    }

    .room--ctrl {
      position: absolute;
      top: 0;
      right: 0;
    }
  }
}

.room-enter-active, .room-leave-active {
  position: absolute;
  transition: transform .1s, opacity .1s;
}
.room-move {
  transition: transform .1s;
}
.room-enter {
  opacity: 0;
  transform: translateY(-50px);
}
.room-leave-to {
  opacity: 0;
  transform: scale(0.8);
}
.room-leave-active {
  position: absolute;
}
</style>

