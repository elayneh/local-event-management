// import { ref } from "vue";
// import { useAuthStore } from "~/stores";
// import { useRouter } from "vue-router";
// import { InsertBoolmarks } from "~/graphql/mutations/event_bookmarks/insert.gql";
// import { DeleteBookmarks } from "~/graphql/mutations/event_bookmarks/delete.gql";

// export const useToggleBookmark = (event) => {
//   const authStore = useAuthStore();
//   const user_id = authStore.id;
//   const router = useRouter();
//   const bookmark = ref(
//     event?.events_bookmarks?.some((bookmark) => bookmark?.user_id === user_id)
//   );

//   const toggleBookmark = async () => {
//     if (!user_id) {
//       return router.push("/users/login");
//     }
//     const isCurrentlyBookmarked = bookmark.value;
//     try {
//       if (!isCurrentlyBookmarked) {
//         await InsertBoolmarks({ user_id, event_id: event.id });
//         bookmark.value = true;
//       } else {
//         await DeleteBookmarks({ user_id, event_id: event.id });
//         bookmark.value = false;
//       }
//     } catch (error) {
//       console.error(
//         `Error ${
//           isCurrentlyBookmarked ? "unbookmarking" : "bookmarking"
//         } event:`,
//         error
//       );
//     }
//   };

//   return { bookmark, toggleBookmark };
// };
