export namespace ytchat {
	
	export class CountData {
	    time: number;
	    count: number;
	
	    static createFrom(source: any = {}) {
	        return new CountData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.time = source["time"];
	        this.count = source["count"];
	    }
	}

}

export namespace ytchop {
	
	export class ChopData {
	    start: string;
	    end: string;
	
	    static createFrom(source: any = {}) {
	        return new ChopData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.start = source["start"];
	        this.end = source["end"];
	    }
	}

}

