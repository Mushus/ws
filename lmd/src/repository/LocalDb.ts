const dbName = 'lmdDB';

export default class LocalDb {
	db: IDBOpenDBRequest;
	constructor() {
		this.db = indexedDB.open(dbName);
		this.db.onupgradeneeded = e => {
			this.db.createObjectStore('files', { keyPath: 'id' });
		};
	}
}
