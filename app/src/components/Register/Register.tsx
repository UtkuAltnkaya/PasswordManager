import { useMutation } from '@tanstack/react-query';
import Loader from 'components/Loader/Loader';
import { MainContext } from 'context/provider';
import { FormEvent, useContext, useState } from 'react';
import { RequestToServer } from 'services/Api';
import style from './Register.module.scss';

interface RegisterInput {
  username: string;
  email: string;
  password: string;
}

const Register = () => {
  const [item, setItem] = useState<RegisterInput>({
    email: '',
    password: '',
    username: '',
  });
  const [passwordConfirm, setPasswordConfirm] = useState('');
  const { setUser } = useContext(MainContext);

  const data = useMutation((item: any) => RequestToServer({ url: 'login', method: 'POST', data: item }), {
    onSuccess: () => {
      setUser(true);
    },
  });

  const handleRegister = (e: FormEvent) => {
    e.preventDefault();
    if (item.password === passwordConfirm) {
      //TODO Error
      return;
    }
    data.mutate({ item });
  };

  return (
    <form className={style.form} onSubmit={handleRegister}>
      {data.isLoading ? (
        <Loader />
      ) : (
        <>
          <div className={style.title}>Register</div>
          <div className={style.formItem}>
            <input
              className={style.input}
              id="username"
              type="text"
              placeholder=" "
              value={item.username}
              onChange={(e) => setItem((item) => ({ ...item, username: e.target.value }))}
            />
            <label htmlFor="username" className={style.label}>
              Username
            </label>
          </div>
          <div className={style.formItem}>
            <input
              className={style.input}
              id="email"
              type="text"
              placeholder=" "
              value={item.email}
              onChange={(e) => setItem((item) => ({ ...item, email: e.target.value }))}
            />
            <label htmlFor="email" className={style.label}>
              Email
            </label>
          </div>
          <div className={style.formItem}>
            <input
              className={style.input}
              id="password"
              type="password"
              placeholder=" "
              value={item.password}
              onChange={(e) => setItem((item) => ({ ...item, password: e.target.value }))}
            />
            <label htmlFor="password" className={style.label}>
              Password
            </label>
          </div>
          <div className={style.formItem}>
            <input
              className={style.input}
              id="passwordConfirm"
              type="password"
              placeholder=" "
              value={passwordConfirm}
              onChange={(e) => setPasswordConfirm(e.target.value)}
            />
            <label htmlFor="passwordConfirm" className={style.label}>
              Password Confirm
            </label>
          </div>
          <div className={style.buttonBox}>
            <div className={style.button}>
              <button
                className={style.buttonDiv}
                // disabled={email === '' || password === ''}
                type="submit"
              >
                Register
              </button>
            </div>
          </div>
        </>
      )}
    </form>
  );
};

export default Register;
