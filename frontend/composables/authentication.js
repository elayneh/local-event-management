import { useMutation } from "@vue/apollo-composable";

export default function (mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "network-only",
  }));
  return {
    onError,
    mutate,
    loading,
    onDone,
  };
}
