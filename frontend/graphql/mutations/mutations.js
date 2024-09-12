import gql from "graphql-tag";

export const DeleteFollowers = gql`
  mutation DeleteFollowers($user_id: uuid!, $event_id: uuid!) {
    delete_followers(
      where: { event_id: { _eq: $event_id }, user_id: { _eq: $user_id } }
    ) {
      returning {
        id
      }
    }
  }
`;

export const InsertFollowers = gql`
  mutation InsertFollowers($event_id: uuid!) {
    insert_followers(objects: { event_id: $event_id }) {
      returning {
        event_id
        id
      }
    }
  }
`;

export const DeleteBookmarks = gql`
  mutation DeleteBookmarks($user_id: uuid!, $event_id: uuid!) {
    delete_bookmarks(
      where: { event_id: { _eq: $event_id }, user_id: { _eq: $user_id } }
    ) {
      returning {
        id
      }
    }
  }
`;

export const InsertBookmarks = gql`
  mutation InsertBookmarks($$event_id: uuid!) {
    insert_bookmarks(objects: { event_id: $event_id }) {
      returning {
        event_id
        id
      }
    }
  }
`;

export const AddCategory = gql`
  mutation AddCategory($name: String!, $description: String!) {
    insert_categories(objects: { name: $name, description: $description }) {
      returning {
        id
        name
        description
      }
    }
  }
`;

export const DeleteCategory = gql`
  mutation DeleteCategory($id: uuid!) {
    delete_categories_by_pk(id: $id) {
      description
      id
      name
    }
  }
`;
