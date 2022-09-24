import { useMutation, useQueryClient } from '@tanstack/react-query';
import { FormEvent, useState } from 'react';
import { RequestToServer } from 'services/Api';
import { isTemplateExpression } from 'typescript';
import style from './Main.module.scss';

type Props = {
  color: string;
};

const AddItem = ({ color }: Props) => {
  const queryClient = useQueryClient();

  const [site, setSite] = useState('');
  const [size, setSize] = useState(32);

  const data = useMutation((item: any) => RequestToServer({ url: 'generatePassword', method: 'POST', data: item }), {
    onSuccess: () => {
      queryClient.invalidateQueries(['user']);
    },
  });

  const handleSubmit = (e: FormEvent) => {
    e.preventDefault();

    if (site === '') {
      return;
    }

    data.mutate({ site, length: size });
    setSite('');
    setSize(32);
  };

  return (
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
  );
};

export default AddItem;
