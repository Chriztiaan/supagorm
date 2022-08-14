/* Do not change, this code is generated from Golang structs */


export class Time {


    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Author {
    id: string;
    createdAt: Time;
    updatedAt: Time;
    name: string;
    userId: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.createdAt = this.convertValues(source["createdAt"], Time);
        this.updatedAt = this.convertValues(source["updatedAt"], Time);
        this.name = source["name"];
        this.userId = source["userId"];
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (a.slice) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}