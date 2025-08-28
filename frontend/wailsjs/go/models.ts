export namespace models {
	
	export class Transaction {
	    id: number;
	    description: string;
	    amount: number;
	    direction: string;
	    category: string;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.amount = source["amount"];
	        this.direction = source["direction"];
	        this.category = source["category"];
	        this.created_at = source["created_at"];
	    }
	}

}

