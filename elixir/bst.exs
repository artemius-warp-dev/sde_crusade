defmodule BinarySearchTree do
  @type bst_node :: %{data: any, left: bst_node | nil, right: bst_node | nil}

  @doc """
  Create a new Binary Search Tree with root's value as the given 'data'
  """
  @spec new(any) :: bst_node
  def new(data) do
    %{data: data, left: nil, right: nil}
  end

  @doc """
  Creates and inserts a node with its value as 'data' into the tree.
  """
  @spec insert(bst_node, any) :: bst_node
  def insert(nil, data), do: new(data)
  def insert(tree, data) do
    cond do
      data > tree.data -> 
        %{tree | right: insert(tree.right, data)}
      data <= tree.data ->
        %{tree | left: insert(tree.left, data)}
    end
  end
  

  @doc """
  Traverses the Binary Search Tree in order and returns a list of each node's data.
  """
  @spec in_order(bst_node) :: [any]
  def in_order(tree) do
    in_order(tree, [])
    |> Enum.reverse()
  end
  defp in_order(nil, acc), do: acc
  defp in_order(tree, acc) do
    left_subtree = in_order(tree.left, acc) 
    acc = [tree.data | left_subtree] 
    in_order(tree.right, acc)
  end
end
