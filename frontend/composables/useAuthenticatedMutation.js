export default function useAuthenticatedMutation (mutation) {
  const { mutate, onDone, loading, onError } = useMutation(mutation, () => ({
    fetchPolicy: "network-only",
    context: {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`, // Get the token from wherever it is stored
        // "x-hasura-role": "user",
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