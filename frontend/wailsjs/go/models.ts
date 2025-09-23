export namespace models {
	
	export class Budget {
	    id: number;
	    name: string;
	    period_start: string;
	    period_end: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Budget(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.period_start = source["period_start"];
	        this.period_end = source["period_end"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Category {
	    id: number;
	    budget_id: number;
	    group_id?: number;
	    name: string;
	    description: string;
	    is_deposit: boolean;
	    expected: number;
	    actual: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.budget_id = source["budget_id"];
	        this.group_id = source["group_id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.is_deposit = source["is_deposit"];
	        this.expected = source["expected"];
	        this.actual = source["actual"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Group {
	    id: number;
	    budget_id: number;
	    name: string;
	    description: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Group(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.budget_id = source["budget_id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Transaction {
	    id: number;
	    description: string;
	    budget_id: number;
	    group_id?: number;
	    category_id?: number;
	    category_name: string;
	    date: string;
	    amount: number;
	    tags: string;
	    notes: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Transaction(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.description = source["description"];
	        this.budget_id = source["budget_id"];
	        this.group_id = source["group_id"];
	        this.category_id = source["category_id"];
	        this.category_name = source["category_name"];
	        this.date = source["date"];
	        this.amount = source["amount"];
	        this.tags = source["tags"];
	        this.notes = source["notes"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}

}

