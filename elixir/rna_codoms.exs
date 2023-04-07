defmodule ProteinTranslation do
  @doc """
  Given an RNA string, return a list of proteins specified by codons, in order.
  """
  @spec of_rna(String.t()) :: {:ok, list(String.t())} | {:error, String.t()}
  def of_rna(rna) do
    rna
    |> to_charlist()
    |> Enum.chunk_every(3)
    |> Enum.map(&List.to_string/1)
    |> Enum.reduce_while([], &construct_rna/2)
    |> wrap_result()
  end

  defp construct_rna(codon, acc) do
    case of_codon(codon) do
      {:ok, "STOP"} -> {:halt, acc}
      {:ok, protein} -> {:cont, [protein | acc]}
      {:error, _} -> {:halt, :error}
    end
  end

  defp wrap_result(proteins) when is_list(proteins), do: {:ok, Enum.reverse(proteins)}
  defp wrap_result(:error), do: {:error, "invalid RNA"}


  @doc """
  Given a codon, return the corresponding protein

  UGU -> Cysteine
  UGC -> Cysteine
  UUA -> Leucine
  UUG -> Leucine
  AUG -> Methionine
  UUU -> Phenylalanine
  UUC -> Phenylalanine
  UCU -> Serine
  UCC -> Serine
  UCA -> Serine
  UCG -> Serine
  UGG -> Tryptophan
  UAU -> Tyrosine
  UAC -> Tyrosine
  UAA -> STOP
  UAG -> STOP
  UGA -> STOP
  """
  @spec of_codon(String.t()) :: {:ok, String.t()} | {:error, String.t()}
  def of_codon(codon) do
    cond do
      codon in ["UGU", "UGC"] -> {:ok, "Cysteine"}
      codon in ["UUA", "UUG"]    -> {:ok, "Leucine"}
      codon in ["AUG"]    -> {:ok, "Methionine"}
      codon in ["UUU", "UUC"] -> {:ok,"Phenylalanine"}
      codon in ["UCU",  "UCC", "UCA", "UCG"] -> {:ok, "Serine"} 
      codon in ["UAU", "UAC"]  -> {:ok, "Tyrosine"}
      codon in ["UGG"] -> {:ok, "Tryptophan"}
      codon in [ "UAA", "UAG", "UGA"] -> {:ok, "STOP"}
      true -> {:error, "invalid codon"}
    end
  end
end
