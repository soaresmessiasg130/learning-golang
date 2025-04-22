import {useEffect, useState} from 'react'
import axios from '../../services/axios/'
import './styles.css'

export default function AvailableAlbums() {
  const [data, setData] = useState([])

  useEffect(() => {
    getData()

    return () => {
      setData([])
    }
  }, [])

  async function getData() {
    try {
      const { data } = await axios.get('/albums')

      setData(data);
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <>
      <h2>Available Albums</h2>

      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>Title</th>
            <th>Artist</th>
            <th>Price</th>
          </tr>
        </thead>

        <tbody>
          {data.map((item, index) => {
            const rowId = 'music-album-list-row-' + index

            return (
              <tr id={rowId} key={rowId}>
                <td>{item.id}</td>
                <td>{item.title}</td>
                <td>{item.artist}</td>
                <td>{item.price}</td>
              </tr>
            )
          })}
        </tbody>
      </table>
    </>
  )
}
