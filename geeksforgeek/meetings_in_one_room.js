//User function Template for javascript

/**
 * @param {number[]} start
 * @param {number[]} end
 * @param {number} n
 * @returns {number}
*/

class Meeting {
    constructor(s,e,i) {
        this.start = s
        this.end = e
        this.index = i
    }
}

class Solution 
{
    //Function to find the maximum number of meetings that can
    //be performed in a meeting room.
    maxMeetings(start, end, n)
    {
        // code here
        let meetings = []
        
        for(let i=0; i<n; i++) {
            meetings.push(new Meeting(start[i], end[i], i+1))
        }
        
        meetings.sort((a,b) => a.end - b.end)
        let _end = meetings[0].end
        let meeting_count = 1
        //console.log(meetings)
        
        for(let i = 1; i<n; i++) {
            if(meetings[i].start > _end) {
                meeting_count++
                _end = meetings[i].end
            }
        }
        return meeting_count
    }
}
