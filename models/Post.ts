/* Do not change, this code is generated from Golang structs */


export class Comment {
    id: string;
    createdAt: Time;
    updatedAt: Time;
    content: string;
    postId: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.createdAt = this.convertValues(source["createdAt"], Time);
        this.updatedAt = this.convertValues(source["updatedAt"], Time);
        this.content = source["content"];
        this.postId = source["postId"];
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
export class Time {


    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Post {
    id: string;
    createdAt: Time;
    updatedAt: Time;
    title: string;
    content: string;
    comments: Comment[];

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.createdAt = this.convertValues(source["createdAt"], Time);
        this.updatedAt = this.convertValues(source["updatedAt"], Time);
        this.title = source["title"];
        this.content = source["content"];
        this.comments = this.convertValues(source["comments"], Comment);
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