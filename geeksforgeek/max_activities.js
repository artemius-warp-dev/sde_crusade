//User function Template for javascript

/**
 * @param {number[]} start
 * @param {number[]} end
 * @param {number} n
 * @returns {number}
*/

class Activity {
    constructor(s,e) {
        this.start = s
        this.end = e
    }
}

class Solution 
{
    //Function to find the maximum number of activities that can
    //be performed by a single person.
    activitySelection(start, end, n)
    {
        // code here
        let activities = []
        let max_activities = 0
        
        for(let i =0; i<n; i++) {
            activities.push(new Activity(start[i], end[i]))
        }
        activities.sort((a,b) => a.end - b.end)
 
        let _start = activities[0].start
        let _end = activities[0].end
        max_activities++
        for(let i=1; i<n; i++) {
            //console.log(activities[i])
            if(activities[i].start > _end) {
                max_activities++
                _end = activities[i].end
            }
        }
        return max_activities
    }
}
