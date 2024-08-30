import * as Yup from "yup";

export const profileValidationSchema = Yup.object({
  username: Yup.string()
    .required("Username is required")
    .min(3, "Username must be at least 3 characters"),
  email: Yup.string()
    .required("Email is required")
    .email("Invalid email format"),
  role: Yup.string().required("Role is required"),
});
