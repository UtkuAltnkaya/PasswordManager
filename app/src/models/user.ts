import type Password from './password';

interface User {
  id: string;
  email: string;
  username: string;
  password_list: Password[];
}

export default User;
