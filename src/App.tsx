import { useEffect, useState } from 'react';
import './App.css';
import { socket } from './socket/socket';

function App() {
  const [initStatus, setInitStatus] = useState<{ id: number; status: string }[]>();

  useEffect(() => {
    console.log(initStatus);
    // Connection opened
    socket.addEventListener('open', () => {
      socket.send('Connection established');
    });

    function handle(event: MessageEvent<string>): void {
      const message: { type: 'init' | 'update'; data: { id: number; status: string }[] } =
        JSON.parse(event.data);
      console.log(message);
      if (message.type === 'init') {
        setInitStatus(message.data);
      } else {
        message.data.forEach((e) => {
          setInitStatus(
            initStatus?.map((i) => {
              if (e.id === i.id) {
                return e;
              } else {
                return i;
              }
            })
          );
        });
      }
    }

    // Listen for messages
    socket.addEventListener('message', handle);

    // setTimeout(() => {
    //   console.log('close');
    //   socket.close();
    // }, 7000);

    // connection.current = ws;

    return () => {
      socket.removeEventListener('message', handle);
    };
  }, [initStatus]);

  return (
    <div>
      {initStatus &&
        initStatus.map((e) => (
          <div key={e.id} style={{ display: 'flex', gap: '10px' }}>
            <p>{e.id}</p>
            <p>{e.status === 'open' ? 'live' : 'not live'}</p>
          </div>
        ))}
      {/* {status &&
        status.map((e) => (
          <div key={e.id} style={{ display: 'flex', gap: '10px' }}>
            <p>{e.id}</p>
            <p>{e.status === 'open' ? 'live' : 'not live'}</p>
          </div>
        ))} */}
    </div>
  );
}

export default App;
