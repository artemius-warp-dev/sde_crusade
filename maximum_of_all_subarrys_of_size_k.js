//User function Template for javascript


class Node {
    constructor() {
        this.data = ""
        this.prev = null
        this.next = null
        
    }
}

class Deque {
    constructor() {
        this.head = null
        this.tail = this.head
    }
    
    append_node(x) {
        let temp = new Node()
        temp.data = x
        temp.prev = temp.next = null

        if(this.head == null) {
            this.head = this.tail = temp
            return 
        }

        this.tail.next = temp
        temp.prev = this.tail
        this.tail = temp
        
    }

    remove_node(node) {
        if(this.head == node)
            this.head = this.head.next
        if(this.tail == node) {
            this.tail = this.tail.prev
            
        }
        if(node.next != null)
            node.next.prev = node.prev
        if(node.prev != null)
            node.prev.next = node.next

    }

    isEmpty() {
        return this.head == null
    }
}



/**
 * @param {number[]} arr
 * @param {number} n
 * @param {number} k
 * @returns {number[]}
*/
class Solution 
{
    //Function to find maximum of each subarray of size k.
    max_of_subarrays(arr, n, k)
    {
        // code here
        let d = new Deque()
        let res = []
        for(let i = 0; i<n; i++) {
            if(d.head && d.head.data == i-k)
                d.remove_node(d.head)

          
            while(d.head && arr[d.tail.data] < arr[i]) {
                d.remove_node(d.tail)
            }
            d.append_node(i)
               
            

            if(i >= k -1) {
                res.push(arr[d.head.data])
            }
        }
        return res
    }
}



let s = new Solution()
let K = 123
let arr = [1909, 1209, 1758, 1221, 1588, 1422, 1946, 1506, 1030, 1413, 1168, 1900, 1591, 1762, 1655, 1410, 1359, 1624, 1537, 1548, 1483, 1595, 1041, 1602, 1350, 1291, 1836, 1374, 1020, 1596, 1021, 1348, 1199, 1668, 1484, 1281, 1734, 1053, 1999, 1418, 1938, 1900, 1788, 1127, 1467, 1728, 1893, 1648, 1483, 1807, 1421, 1310, 1617, 1813, 1514, 1309, 1616, 1935, 1451, 1600, 1249, 1519, 1556, 1798, 1303, 1224, 1008, 1844, 1609, 1989, 1702, 1195, 1485, 1093, 1343, 1523, 1587, 1314, 1503, 1448, 1200, 1458, 1618, 1580, 1796, 1798, 1281, 1589, 1798, 1009, 1157, 1472, 1622, 1538, 1292, 1038, 1179, 1190, 1657, 1958, 1191, 1815, 1888, 1156, 1511, 1202, 1634, 1272, 1055, 1328, 1646, 1362, 1886, 1875, 1433, 1869, 1142, 1844, 1416, 1881, 1998, 1322, 1651, 1021, 1699, 1557, 1476, 1892, 1389, 1075, 1712, 1600, 1510, 1003, 1869, 1861, 1688, 1401, 1789, 1255, 1423, 1002, 1585, 1182, 1285, 1088, 1426, 1617, 1757, 1832, 1932, 1169, 1154, 1721, 1189, 1976, 1329, 1368, 1692, 1425, 1555, 1434, 1549, 1441, 1512, 1145, 1060, 1718, 1753, 1139, 1423, 1279, 1996, 1687, 1529, 1549, 1437, 1866, 1949, 1193, 1195, 1297, 1416]
let N = arr.length
console.log(s.max_of_subarrays(arr, N, K))
