import { useQuery, useMutation } from "@vue/apollo-composable";
import getSingleTag from "~/graphql/queries/tags/getSingleTag.gql";
import UpdateTagsByPk from "~/graphql/mutations/tags/update.gql";

export default function useTag() {
  const fetchTagById = async (id) => {
    const { result, loading, error } = useQuery(
      getSingleTag,
      { id },
      {
        fetchPolicy: "no-cache",
        clientId: "authClient",
        context: {
          headers: {
            "x-hasura-role": "admin",
          },
        },
      }
    );

    if (loading) {
      console.log("Loading tag...");
      return null;
    }

    if (error) {
      console.error("Error fetching tag:", error);
      return null;
    }

    return result.value?.tags_by_pk || null;
  };

  const updateTag = async (id, values) => {
    const { mutate } = useMutation(UpdateTagsByPk, {
      fetchPolicy: "no-cache",
      clientId: "authClient",
      context: {
        headers: {
          "x-hasura-role": "admin",
        },
      },
    });

    try {
      const { data } = await mutate({ id, ...values });
      return !!data;
    } catch (error) {
      console.error("Error updating tag:", error);
      return false;
    }
  };

  return { fetchTagById, updateTag };
}
