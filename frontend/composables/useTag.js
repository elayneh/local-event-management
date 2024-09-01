import { useMutation } from "@vue/apollo-composable";
import getSingleTag from "~/graphql/queries/tags/getSingleTag.gql";
import UpdateTagsByPk from "~/graphql/mutations/tags/update.gql";

export default function useTag() {
  const fetchTagById = async (id) => {
    const { mutate: getTagMutation } = useMutation(getSingleTag);
    const { data } = await getTagMutation({ id });
    return data?.tags_by_pk;
  };

  const updateTag = async (id, values) => {
    const { mutate: updateTagMutation } = useMutation(UpdateTagsByPk);
    const { data } = await updateTagMutation({ id, ...values });
    return !!data;
  };

  return { fetchTagById, updateTag };
}
