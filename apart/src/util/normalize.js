import { MAX_DATE } from '@/util/constants';

export function normalizeArticle(id = null, data = {}) {
  return {
    id,
    data: {
      name: String(data.name || ''),
      administrator: String(data.administrator || ''),
      parkingFee: Number(data.parkingFee || 0),
      commonAreaCharge: Number(data.commonAreaCharge || 0)
    }
  };
}

export function normalizeRoom(id = null, data = {}) {
  return {
    id,
    data: {
      articleId: String(data.articleId || ''),
      index: Number(data.index || 0),
      name: String(data.name || ''),
      rent: Number(data.rent || 0)
    }
  };
}

export function normalizeTenant(id = null, data = {}) {
  return {
    id,
    data: {
      articleId: String(data.articleId || ''),
      commonAreaCharge: Number(data.commonAreaCharge || 0),
      moveInAt: Number(data.moveInAt || 0),
      moveOutAt: Number(data.moveOutAt || MAX_DATE),
      name: String(data.name || ''),
      parkingFee: Number(data.name || 0),
      rent: Number(data.rent || 0),
      roomId: String(data.roomId || '')
    }
  };
}

export function normalizeReceipt(id = null, data = {}) {
  return {
    id,
    data: {
      publishAt: Number(data.publishAt || 0),
      tenantName: String(data.tenantName || ''),
      rent: Number(data.rent || 0),
      commonAreaCharge: Number(data.commonAreaCharge || 0),
      parkingFee: Number(data.parkingFee || 0),
      waterCharge: Number(data.waterCharge || 0),
      administrator: String(data.administrator || ''),
      tenantId: String(data.tenantId),
      date: String(data.date)
    }
  };
}
