import { useMutation, useQueryClient } from '@tanstack/react-query';
import Data from 'models/data';
import Password from 'models/password';
import { FormEvent, useState } from 'react';
import { RequestToServer } from 'services/Api';
import style from './Main.module.scss';
import { toast } from 'react-toastify';

type Props = {
  color: string;
};

const AddItem = ({ color }: Props) => {
  const queryClient = useQueryClient();

  const [site, setSite] = useState('');
  const [size, setSize] = useState(32);

  const data = useMutation<Data<Password>, Data<string>, any>(
    (item: any) => RequestToServer({ url: 'generatePassword', method: 'POST', data: item }),
    {
      onSuccess: () => queryClient.invalidateQueries(['user']),
    },
  );

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    if (site === '') {
      return toast.error('Site must not be empty');
    }
    try {
      await data.mutateAsync({ site, length: size });
      setSite('');
      setSize(32);
      return toast.success('Password Generated for ' + site);
    } catch (error) {
      setSite('');
      setSize(32);
      return toast.error(error.message);
    }
  };

  return (
    <>
      <div className={style.addItem}>
        <form onSubmit={handleSubmit}>
          <input
            className={style.input}
            style={{ backgroundColor: color }}
            type="text"
            placeholder="Add item"
            value={site}
            onChange={(e) => setSite(e.target.value)}
          />
          <select value={size} onChange={(e) => setSize(parseInt(e.target.value))} style={{ backgroundColor: color }}>
            <option value={4}>4</option>
            <option value={8}>8</option>
            <option value={16}>16</option>
            <option value={32}>32</option>
            <option value={64}>64</option>
          </select>
          <button style={{ backgroundColor: color }}>Add</button>
        </form>
      </div>
    </>
  );
};

export default AddItem;
