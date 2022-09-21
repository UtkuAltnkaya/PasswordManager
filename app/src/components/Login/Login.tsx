import { useMutation } from '@tanstack/react-query';
import Loader from 'components/Loader/Loader';
import { MainContext } from 'context/provider';

import { FormEvent, useContext, useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { RequestToServer } from 'services/Api';
import style from './Login.module.scss';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { setUser } = useContext(MainContext);

  const navigate = useNavigate();
  const data = useMutation((item: any) => RequestToServer({ url: 'login', method: 'POST', data: item }), {
    onSuccess: () => {
      setUser(true);
    },
  });
  const handleLogin = (e: FormEvent) => {
    e.preventDefault();

    data.mutate({ email, password });
  };

  useEffect(() => {
    if (data.isSuccess) {
      return navigate('/');
    }
  }, [data.isSuccess, navigate]);

  return (
    <form className={style.form} onSubmit={handleLogin}>
      {data.isLoading ? (
        <div className={style.formItem}>
          <Loader />
        </div>
      ) : (
        <>
          <div className={style.title}>Login</div>
          <div className={style.formItem}>
            <input
              className={style.input}
              id="email"
              type="text"
              placeholder=" "
              value={email}
              onChange={(e) => setEmail(e.target.value)}
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
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <label htmlFor="password" className={style.label}>
              Password
            </label>
          </div>
          <div className={style.buttonBox}>
            <div className={style.button}>
              <button className={style.buttonDiv} disabled={email === '' || password === ''} type="submit">
                Login
              </button>
            </div>
          </div>
        </>
      )}
    </form>
  );
};

export default Login;
