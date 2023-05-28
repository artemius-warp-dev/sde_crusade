defmodule Satellite do
  @typedoc """
  A tree, which can be empty, or made from a left branch, a node and a right branch
  """
  @type tree :: {} | {tree, any, tree}

  @doc """
  Build a tree from the elements given in a pre-order and in-order style
  """
  @spec build_tree(preorder :: [any], inorder :: [any]) :: {:ok, tree} | {:error, String.t()}
  def build_tree([], _), do: {:ok, {}}

  def build_tree(preorder, inorder) do
    cond do
      length(preorder) != length(inorder) ->
        {:error, "traversals must have the same length"}

      preorder -- inorder != [] ->
        {:error, "traversals must have the same elements"}

      length(Enum.uniq(preorder)) != length(inorder) ->
        {:error, "traversals must contain unique items"}

      true ->
        {:ok, do_build_tree(preorder, inorder)}
    end
  end

  defp do_build_tree([], []), do: {}

  defp do_build_tree([root | rest], inorder) do
    {left_in, [_ | right_in]} = inorder |> Enum.split_while(&(&1 != root))

    {left_pre, right_pre} = rest |> Enum.split(Enum.count(left_in))

    {
      do_build_tree(left_pre, left_in),
      root,
      do_build_tree(right_pre, right_in)
    }
  end
end

# Satellite.build_tree(~w(a i c d x f r)a, ~w(c i d a f x r)a) |> IO.inspect()
preorder = ~w(a i x f r)a

inorder = ~w(i a f x r)a

tree = {{{}, :i, {}}, :a, {{{}, :f, {}}, :x, {{}, :r, {}}}}
Satellite.build_tree(preorder, inorder) |> IO.inspect() == {:ok, tree}
