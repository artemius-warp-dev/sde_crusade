// Given a binary matrix of M X N of integers,
// you need to return only unique rows of binary array
let ROW = 3
let COL = 4

// A Trie node
class Node
{
	constructor()
	{
		this.isEndOfCol;
		this.child = new Array(2); // Only two children needed for 0 and 1
	}
} ;


// A utility function to allocate memory
// for a new Trie node
function newNode()
{
	let temp = new Node();
	temp.isEndOfCol = 0;
	temp.child[0] = null;
	temp.child[1] = null;
	return temp;
}

// Inserts a new matrix row to Trie.
// If row is already present,
// then returns 0, otherwise insets the row and
// return 1
function insert(root, M, row, col)
{
	// base case
	if (root == null)
		root = newNode();

	// Recur if there are more entries in this row
	if (col < (M.length).length)
		return insert(((root).child[M[row][col]]),
										M, row, col + 1);

	else // If all entries of this row are processed
	{
		// unique row found, return 1
		if (!((root).isEndOfCol))
		{
			(root).isEndOfCol = 1;
			return 1;
		}

		// duplicate row found, return 0
	}
	return 0;
}

// A utility function to print a row
function printRow(M, row)
{
	console.log(M[row].join(" "))
}

// The main function that prints
// all unique rows in a given matrix.
function findUniqueRows(M)
{
	let root = null; // create an empty Trie
	let i;

	// Iterate through all rows
	for (i = 0; i < ROW; ++i)
	
		// insert row to TRIE
		if (insert(root, M, i, 0))
		
			// unique row found, print it
			printRow(M, i);
}

// Driver Code
let M = [ [ 1, 1, 0, 1 ], [ 1, 0, 0, 1 ], [ 1, 1, 0, 1 ] ]

findUniqueRows(M);

// This code is contributed by phasing17
