export default function useAuthenticatedMutation(mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "network-only",
    context: {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "x-hasura-role": "admin",
      },
    },
  }));

  return {
    mutate,
    onDone,
    loading,
    onError,
  };
}
