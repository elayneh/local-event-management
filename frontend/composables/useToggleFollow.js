import { ref } from "vue";
import { useMutation } from "@vue/apollo-composable";
import {
  InsertFollowers,
  DeleteFollowers,
} from "~/graphql/mutations/event_followers.gql";
import { useAuthStore } from "~/stores";

export function useToggleFollow(event) {
  const following = ref(false);

  const { mutate: insertFollow } = useMutation(InsertFollowers);
  const { mutate: deleteFollow } = useMutation(DeleteFollowers);
  const userId = useAuthStore().id;


  
  const toggleFollow = async () => {
    try {
      if (following.value) {
        await deleteFollow({ eventId: event.id, user_id: userId }); // Replace with actual userId
        following.value = false;
        console.log("Unfollowed event");
      } else {
        await insertFollow({ eventId: event.id, user_id: userId }); // Replace with actual userId
        following.value = true;
        console.log("Followed event");
      }
    } catch (error) {
      console.error(
        `Error ${following.value ? "unfollowing" : "following"} event:`,
        error
      );
    }
  };

  return { following, toggleFollow };
}
