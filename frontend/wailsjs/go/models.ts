export namespace models {
	
	export class Budget {
	    id: number;
	    project_id: number;
	    name: string;
	    period_start: string;
	    period_end: string;
	    expected_income: number;
	    starting_balance: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Budget(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.project_id = source["project_id"];
	        this.name = source["name"];
	        this.period_start = source["period_start"];
	        this.period_end = source["period_end"];
	        this.expected_income = source["expected_income"];
	        this.starting_balance = source["starting_balance"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class BudgetAllocation {
	    id: number;
	    budget_id: number;
	    category_id: number;
	    expected_cost: number;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new BudgetAllocation(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.budget_id = source["budget_id"];
	        this.category_id = source["category_id"];
	        this.expected_cost = source["expected_cost"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Category {
	    id: number;
	    group_id?: number;
	    name: string;
	    description: string;
	    expense_type: boolean;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.group_id = source["group_id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.expense_type = source["expense_type"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Group {
	    id: number;
	    project_id: number;
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
	        this.project_id = source["project_id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Project {
	    id: number;
	    name: string;
	    description: string;
	    currency: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.description = source["description"];
	        this.currency = source["currency"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Tag {
	    id: number;
	    project_id: number;
	    name: string;
	    created_at: string;
	    updated_at: string;
	
	    static createFrom(source: any = {}) {
	        return new Tag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.project_id = source["project_id"];
	        this.name = source["name"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class Transaction {
	    id: number;
	    description: string;
	    project_id: number;
	    category_id?: number;
	    category_name: string;
	    date: string;
	    amount: number;
	    expense_type: boolean;
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
	        this.project_id = source["project_id"];
	        this.category_id = source["category_id"];
	        this.category_name = source["category_name"];
	        this.date = source["date"];
	        this.amount = source["amount"];
	        this.expense_type = source["expense_type"];
	        this.notes = source["notes"];
	        this.created_at = source["created_at"];
	        this.updated_at = source["updated_at"];
	    }
	}
	export class TransactionTag {
	    transaction_id: number;
	    tag_id: number;
	    created_at: string;
	
	    static createFrom(source: any = {}) {
	        return new TransactionTag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.transaction_id = source["transaction_id"];
	        this.tag_id = source["tag_id"];
	        this.created_at = source["created_at"];
	    }
	}

}

