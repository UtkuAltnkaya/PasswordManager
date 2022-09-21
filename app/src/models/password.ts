interface Password {
  id: string;
  password: string;
  password_name: PasswordName;
  password_name_id: string;
  userId: string;
}

interface PasswordName {
  id: string;
  name: string;
}

export default Password;
