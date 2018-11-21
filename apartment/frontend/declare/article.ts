export interface Article {
  id: number;
  name: string;
}

export interface ArticleDetail extends Article {
  rooms: Room[];
}

export interface Room {
  id: number;
  name: string;
  rent: number;
  index: number;
}

export interface RoomDetail extends Room {
  tenants: Tenant[];
}

export interface Tenant {
  id: number;
  name: string;
  rent: number;
  moveInAt: string;
  moveOutAt: string;
}

export interface SearchTenantArticle {
  id: number;
  name: string;
  rooms: Room[];
}

export interface SearchTenantTenant {
  id: number;
  name: string;
  index: number;
  tenants: SearchTenant[];
}

export interface SearchTenant {
  id: number;
  name: string;
}

export enum RoomStatus {
  Empty = "empty",
  Booking = "booking",
  Living = "living"
}

export interface RoomShowStatus {
  id: number;
  name: string;
  rent: number;
  status: RoomStatus;
  article: Article;
  tenants: Tenant[];
}

export interface ListBill {
  articles: ListBillArticle[];
}

export interface ListBillArticle {
  id: number;
  name: string;
  rooms: ListBillRoom[];
}

export interface ListBillRoom {
  id: number;
  name: string;
  index: number;
  teant: ListBillTenant
}

export interface ListBillTenant {
  id: number;
  name: string;
  rent: number;
  issueFlag: boolean;
}
